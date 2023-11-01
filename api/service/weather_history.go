package service

import (
	"time"
	"weatherapp/db/weather_history"
	"weatherapp/pkg/logging"
)

type WeatherHistory interface {
	CreateRecord(userID int, cityName string, weatherData string) error
	GetRecordByID(id int) (*weather_history.Record, error)
	GetPaginatedRecords(userID int, recordLen int, index int) ([]*weather_history.Record, error)
	DeleteRecord(id int) error
	BulkDeleteRecords(recordIDs []int) error
}

//go:generate mockgen -package=service -source=weather_history.go -destination=weather_history_mock.go

type weatherHistory struct {
	store  weather_history.Store
	logger logging.Logger
}

// NewService creates a new instance of the Weather History service.
func NewWeatherHistoryService(store weather_history.Store, logger logging.Logger) WeatherHistory {
	return &weatherHistory{store: store, logger: logger}
}

// CreateRecord creates a new weather history record.
func (s *weatherHistory) CreateRecord(userID int, cityName string, weatherData string) error {
	record := &weather_history.Record{
		UserID:      userID,
		CityName:    cityName,
		WeatherData: weatherData,
		SearchTime:  time.Now(),
	}

	return s.store.CreateRecord(record)
}

// GetRecordByID retrieves a weather history record by its ID.
func (s *weatherHistory) GetRecordByID(id int) (*weather_history.Record, error) {
	return s.store.GetRecordByID(id)
}

// GetPaginatedRecords retrieves paginated weather history records for a user.
func (s *weatherHistory) GetPaginatedRecords(userID int, recordLen int, index int) ([]*weather_history.Record, error) {
	return s.store.GetPaginatedRecords(userID, recordLen, index)
}

// DeleteRecord deletes a weather history record by its ID.
func (s *weatherHistory) DeleteRecord(id int) error {
	return s.store.DeleteRecord(id)
}

// BulkDeleteRecords deletes multiple weather history records.
func (s *weatherHistory) BulkDeleteRecords(recordIDs []int) error {
	return s.store.BulkDeleteRecords(recordIDs)
}
