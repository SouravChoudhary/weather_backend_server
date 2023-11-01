package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"weatherapp/api/service"
	"weatherapp/db/weather_history"
	"weatherapp/models"
	"weatherapp/pkg/logging"
)

type WeatherHistoryHandler interface {
	CreateRecord(w http.ResponseWriter, r *http.Request)
	GetRecordByID(w http.ResponseWriter, r *http.Request)
	GetPaginatedRecords(w http.ResponseWriter, r *http.Request)
	DeleteRecord(w http.ResponseWriter, r *http.Request)
	BulkDeleteRecords(w http.ResponseWriter, r *http.Request)
}

type weatherHistory struct {
	service service.WeatherHistory
	logger  logging.Logger
}

func NewWeatherHistoryHandler(weatherHistoryService service.WeatherHistory, logger logging.Logger) WeatherHistoryHandler {
	return &weatherHistory{
		service: weatherHistoryService,
		logger:  logger,
	}
}

func (api *weatherHistory) CreateRecord(w http.ResponseWriter, r *http.Request) {
	var record weather_history.Record
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Error: "Invalid request data"})
		return
	}

	if err := api.service.CreateRecord(record.UserID, record.CityName, record.WeatherData); err != nil {
		respondWithJSON(w, http.StatusInternalServerError, models.Response{Error: err.Error()})
		return
	}
	respondWithJSON(w, http.StatusOK, models.Response{Data: "Weather record created successfully"})
}

func (api *weatherHistory) GetRecordByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Error: "Invalid ID"})
		return
	}

	record, err := api.service.GetRecordByID(id)
	if err != nil {
		respondWithJSON(w, http.StatusNotFound, models.Response{Error: "Record not found"})
		return
	}

	respondWithJSON(w, http.StatusOK, models.Response{Data: record})
}

func (api *weatherHistory) GetPaginatedRecords(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	perPageStr := r.URL.Query().Get("per_page")
	userIDStr := r.URL.Query().Get("user_id")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Error: "Invalid page parameter"})
		return
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage <= 0 {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Error: "Invalid per_page parameter"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil || page <= 0 {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Error: "Invalid user id parameter"})
		return
	}

	records, err := api.service.GetPaginatedRecords(userID, perPage, page)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, models.Response{Error: "Failed to fetch paginated records"})
		return
	}
	respondWithJSON(w, http.StatusOK, models.Response{Data: records})
}

func (api *weatherHistory) DeleteRecord(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Error: "Invalid ID"})
		return
	}

	if err := api.service.DeleteRecord(id); err != nil {
		respondWithJSON(w, http.StatusInternalServerError, models.Response{Error: "Failed to delete weather record"})
		return
	}
	respondWithJSON(w, http.StatusOK, models.Response{Data: "Weather record deleted successfully"})
}

func (api *weatherHistory) BulkDeleteRecords(w http.ResponseWriter, r *http.Request) {
	var request models.BulkDeleteRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithJSON(w, http.StatusBadRequest, models.Response{Error: "Invalid request data"})
		return
	}

	if err := api.service.BulkDeleteRecords(request.RecordIDs); err != nil {
		respondWithJSON(w, http.StatusInternalServerError, models.Response{Error: "Failed to bulk delete weather records"})
		return
	}
	respondWithJSON(w, http.StatusOK, models.Response{Data: "Weather records deleted successfully"})
}
