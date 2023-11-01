package handlers

import (
	"net/http"
	"strconv"
	"weatherapp/api/service"
	"weatherapp/models"
	"weatherapp/pkg/logging"
)

type WeatherApiHandler interface {
	GetCurrentWeather(w http.ResponseWriter, r *http.Request)
}

type WeatherAPI struct {
	weatherService service.WeatherService
	logger         logging.Logger
}

func NewWeatherAPI(weatherService service.WeatherService, logger logging.Logger) WeatherApiHandler {
	return &WeatherAPI{
		weatherService: weatherService,
		logger:         logger,
	}
}

func (api *WeatherAPI) GetCurrentWeather(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	cityName := r.URL.Query().Get("city")

	if cityName == "" {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Error: "City name is required"})
	}

	weatherData, err := api.weatherService.GetCurrentWeather(userID, cityName)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, models.Response{Error: err.Error()})
		return
	}

	respondWithJSON(w, http.StatusOK, models.Response{Data: weatherData})
}
