package config

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type App struct {
	Logger               Logger               `mapstructure:"logger"`
	JWTAuthenticator     JWTAuthenticator     `mapstructure:"jwt_authenticator"`
	OpenWeatherMapClient OpenWeatherMapClient `mapstructure:"open_weather_map_client"`
	Database             Database             `mapstructure:"database"`
}

type Database struct {
	DSN string `mapstructure:"dsn"`
}

type Logger struct {
	Type     string `mapstructure:"type"`
	Output   string `mapstructure:"output"`
	FileName string `mapstructure:"file_name"`
}

type JWTAuthenticator struct {
	SecretKey string        `mapstructure:"secret_key"`
	TokenTTL  time.Duration `mapstructure:"token_ttl"`
}

type OpenWeatherMapClient struct {
	APIKey          string        `mapstructure:"api_key"`
	Hostname        string        `mapstructure:"hostname"`
	BaseURL         string        `mapstructure:"base_url"`
	Timeout         time.Duration `mapstructure:"timeout"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	IdleConnTimeout time.Duration `mapstructure:"idle_conn_timeout"`
}

// Load configuration from environment variables or configuration files
func LoadConfig() (*App, error) {
	configPath := getConfigPath()
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config:LoadConfig: configuration file not found")
		}
		return nil, errors.New("config:LoadConfig: failed to read configuration: err: " + err.Error())
	}

	var cfg App
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.New("config:LoadConfig: failed to unmarshal configuration: err: " + err.Error())
	}
	if err := validateConfig(cfg); err != nil {
		return nil, errors.New("config:LoadConfig: configuration validation failed: err: " + err.Error())
	}

	return &cfg, nil
}

func validateConfig(cfg App) error {
	if cfg.Database.DSN == "" {
		return errors.New("database DSN is required")
	}
	if cfg.JWTAuthenticator.SecretKey == "" {
		return errors.New("JWT secret key is required")
	}
	if cfg.OpenWeatherMapClient.APIKey == "" {
		return errors.New("OpenWeatherMap API key is required")
	}
	if cfg.OpenWeatherMapClient.Hostname == "" {
		return errors.New("OpenWeatherMapClient hostname is required")
	}

	return nil
}

// getConfigPath returns the absolute file path to the configuration directory.
// It assumes that the service is started from the root directory (e.g., using "go run ./cmd/server/main.go").
// The function constructs the configuration path by appending the relative path to the root directory
// and replaces backslashes with forward slashes for cross-platform compatibility.
func getConfigPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configPath := cwd + `\config`
	configPath = strings.ReplaceAll(configPath, `\`, "/")
	return configPath
}
