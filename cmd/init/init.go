// init/init.go
package init

import (
	"errors"
	"weatherapp/api/handlers"
	"weatherapp/api/service"
	"weatherapp/config"
	"weatherapp/db"
	"weatherapp/db/user"
	"weatherapp/db/weather_history"
	"weatherapp/external"
	"weatherapp/pkg/auth"
	"weatherapp/pkg/logging"
)

// InitializeApp initializes and returns app handlers.
func InitializeApp(appConfig *config.App, appLogger logging.Logger, jwtAuthenticator auth.JWTAuthenticator) (handlers.AuthHandler,
	handlers.WeatherHistoryHandler,
	handlers.WeatherApiHandler, error) {

	// Initialize the database connection
	dbConnection, err := db.ConnectToDatabase(appConfig.Database.DSN)
	if err != nil {
		return nil, nil, nil, errors.New("init: InitializeApp: Failed to connect to the database: err: " + err.Error())
	}

	// create schema
	dbConnection.AutoMigrate(weather_history.Record{})
	dbConnection.AutoMigrate(user.User{})

	// Create stores
	weatherHistoryStore := weather_history.NewStore(dbConnection, appLogger)
	userStore := user.NewStore(dbConnection, appLogger)

	// open weather client
	openWeatherClient := external.NewOpenWeatherMapClient(appConfig.OpenWeatherMapClient, appLogger)

	// Create services
	authService := service.NewAuthService(userStore, jwtAuthenticator, appLogger)
	weatherHistoryService := service.NewWeatherHistoryService(weatherHistoryStore, appLogger)
	currentweatherService := service.NewWeatherService(weatherHistoryService, openWeatherClient, appLogger)

	// Create and configure handlers
	authHandler := handlers.NewAuthHandler(authService, appLogger)
	weatherHistoryHandler := handlers.NewWeatherHistoryHandler(weatherHistoryService, appLogger)
	currentWeatherHandler := handlers.NewWeatherAPI(currentweatherService, appLogger)

	return authHandler,
		weatherHistoryHandler,
		currentWeatherHandler, nil
}
