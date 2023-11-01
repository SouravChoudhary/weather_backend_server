package service

import (
	"weatherapp/apperror"
	"weatherapp/external"
	"weatherapp/pkg/logging"
)

type WeatherService interface {
	GetCurrentWeather(userID int, cityName string) (string, error)
}

// WeatherService implements the WeatherService interface.
type weatherService struct {
	weatherHistorySvc WeatherHistory
	weatherAPIClient  external.OpenWeatherMapClient
	logger            logging.Logger
}

// NewWeatherService creates a new WeatherService instance with the provided dependencies.
func NewWeatherService(weatherHistorySvc WeatherHistory, weatherAPIClient external.OpenWeatherMapClient, logger logging.Logger) WeatherService {
	return &weatherService{
		weatherHistorySvc: weatherHistorySvc,
		weatherAPIClient:  weatherAPIClient,
		logger:            logger,
	}
}

// GetCurrentWeather retrieves the current weather for the specified city.
func (ws *weatherService) GetCurrentWeather(userID int, cityName string) (string, error) {

	weatherData, err := ws.weatherAPIClient.GetWeatherData(cityName)
	if err != nil {
		return "", apperror.ReturnServiceErr("WC01", err)
	}

	err = ws.weatherHistorySvc.CreateRecord(userID, cityName, weatherData)
	if err != nil {
		return "", err
	}

	return weatherData, nil
}
