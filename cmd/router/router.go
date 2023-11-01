package router

import (
	"net/http"
	"weatherapp/api/middleware"
	appServiceInit "weatherapp/cmd/init"
	"weatherapp/config"
	"weatherapp/pkg/auth"
	"weatherapp/pkg/logging"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SetupRouter(appConfig *config.App, appLogger logging.Logger) (http.Handler, error) {

	jwtAuthenticator := auth.NewJWTAuthenticator(appConfig.JWTAuthenticator, appLogger)

	// Initialize the app components using the init package
	authHandler,
		weatherHistoryHandler,
		currentWeatherHandler, err := appServiceInit.InitializeApp(
		appConfig,
		appLogger,
		jwtAuthenticator,
	)

	if err != nil {
		return nil, err
	}

	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(r)

	// Additional configurations
	r.Use(middleware.LoggingMiddleware(appLogger)) // Add request logging middleware

	// Apply token authentication middleware to protected routes
	protectedRouter := r.PathPrefix("/protected").Subrouter()
	protectedRouter.Use(middleware.TokenAuthMiddleware(jwtAuthenticator))

	// protected routes with their corresponding handlers
	protectedRouter.HandleFunc("/weather/current", currentWeatherHandler.GetCurrentWeather).Methods("GET")
	protectedRouter.HandleFunc("/weather/history", weatherHistoryHandler.GetPaginatedRecords).Methods("GET")
	protectedRouter.HandleFunc("/weather/history/bulkdelete", weatherHistoryHandler.BulkDeleteRecords).Methods("POST")

	// login and registration routes
	r.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/auth/register", authHandler.Register).Methods("POST")

	return handler, nil
}
