package logging

import (
	"io"
	gologger "log"
	"os"
	"weatherapp/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger interface {
	Log(level string, message string, fields map[string]interface{})
	LogInfo(message string, fields map[string]interface{})
	LogDebug(message string, fields map[string]interface{})
	LogWarning(message string, fields map[string]interface{})
	LogError(message string, fields map[string]interface{})
	LogFatal(message string, fields map[string]interface{})
	LogPanic(message string, fields map[string]interface{})
	LogTrace(message string, fields map[string]interface{})
	LogAudit(message string, fields map[string]interface{})
	LogSecurity(message string, fields map[string]interface{})
}

//go:generate mockgen -package=logging -source=logger.go -destination=logger_mock.go

type logger struct {
	log zerolog.Logger
}

// Log rotation not implemented
func NewLogger(config config.Logger) Logger {

	var l zerolog.Logger
	if config.FileName == "" {
		config.FileName = "app.log"
	}
	switch config.Type {
	case "zerolog":
		switch config.Output {
		case "stdout":
			l = log.Output(io.MultiWriter(os.Stdout, createLogFile(config.FileName)))
		default:
			l = log.Output(io.MultiWriter(os.Stdout, createLogFile(config.FileName)))
		}
	default:
		l = log.Output(io.MultiWriter(os.Stdout, createLogFile(config.FileName)))
	}
	return &logger{log: l}
}

func (l *logger) logWithFields(level string, message string, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	// Log with the specified log level and structured fields
	logEntry := l.log.With().Fields(fields).Logger()
	logEntry.Log().Str("level", level).Msg(message)
}

func (l *logger) Log(level string, message string, fields map[string]interface{}) {
	l.logWithFields(level, message, fields)
}

func (l *logger) LogInfo(message string, fields map[string]interface{}) {
	l.logWithFields("info", message, fields)
}

func (l *logger) LogDebug(message string, fields map[string]interface{}) {
	l.logWithFields("debug", message, fields)
}

func (l *logger) LogWarning(message string, fields map[string]interface{}) {
	l.logWithFields("warning", message, fields)
}

func (l *logger) LogError(message string, fields map[string]interface{}) {
	l.logWithFields("error", message, fields)
}

func (l *logger) LogFatal(message string, fields map[string]interface{}) {
	l.logWithFields("fatal", message, fields)
}

func (l *logger) LogPanic(message string, fields map[string]interface{}) {
	l.logWithFields("panic", message, fields)
}

func (l *logger) LogTrace(message string, fields map[string]interface{}) {
	l.logWithFields("trace", message, fields)
}

func (l *logger) LogAudit(message string, fields map[string]interface{}) {
	l.logWithFields("audit", message, fields)
}

func (l *logger) LogSecurity(message string, fields map[string]interface{}) {
	l.logWithFields("security", message, fields)
}

func createLogFile(filename string) io.Writer {
	file, err := os.Create(filename)
	if err != nil {
		gologger.Fatalf("logging: createLogFile: Failed to create log file: %v", err)
		return os.Stdout
	}
	return file
}
