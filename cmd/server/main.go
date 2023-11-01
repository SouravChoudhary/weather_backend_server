package main

import (
	"log"
	"net/http"
	"os"
	"weatherapp/cmd/router"
	"weatherapp/config"
	"weatherapp/pkg/logging"
)

func main() {

	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
	}

	appLogger := logging.NewLogger(appConfig.Logger)

	// Create a router and apply middleware
	r, err := router.SetupRouter(appConfig, appLogger)
	if err != nil {
		appLogger.LogFatal(err.Error(), nil)
		os.Exit(1)
	}

	// Start the HTTP server
	http.Handle("/", r)
	appLogger.LogInfo("Application starting at port :8080", nil)
	defer func() {
		appLogger.LogInfo("Application stopped", nil)
	}()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		appLogger.LogError("Application stopped", map[string]interface{}{"error": err.Error()})
	}
}
