package weather_history

import (
	"weatherapp/apperror"
	"weatherapp/pkg/logging"

	"gorm.io/gorm"
)

type Store interface {
	CreateRecord(record *Record) error
	GetRecordByID(id int) (*Record, error)
	GetPaginatedRecords(userID int, recordLen int, index int) ([]*Record, error)
	DeleteRecord(id int) error
	BulkDeleteRecords(recordIDs []int) error
}

type StoreGorm struct {
	db     *gorm.DB
	logger logging.Logger
}

func NewStore(db *gorm.DB, logger logging.Logger) Store {
	return &StoreGorm{db: db, logger: logger}
}

// CreateRecord creates a weather record in the database.
func (wh *StoreGorm) CreateRecord(record *Record) error {
	if err := wh.db.Create(record).Error; err != nil {
		wh.logger.LogError("fail to create weather record", map[string]interface{}{"package": "weather_history", "method": "CreateRecord", "record_id": record.ID, "error": err.Error()})
		return apperror.ReturnStoreErr("WH03", err)
	}
	return nil
}

// GetRecordByID retrieves a weather record by its ID.
func (wh *StoreGorm) GetRecordByID(id int) (*Record, error) {
	var record Record
	if err := wh.db.Where("id = ?", id).First(&record).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperror.WeatherHistoryRecordNotFound
		}
		wh.logger.LogError("fail to get weather record by ID", map[string]interface{}{"package": "weather_history", "method": "GetRecordByID", "record_id": id, "error": err.Error()})
		return nil, apperror.ReturnStoreErr("WH01", err)
	}
	return &record, nil
}

// GetPaginatedRecords retrieves paginated weather records for a specific user.
func (wh *StoreGorm) GetPaginatedRecords(userID int, recordLen int, index int) ([]*Record, error) {
	var records []*Record
	if err := wh.db.Where("user_id = ?", userID).
		Offset((index - 1) * recordLen).
		Limit(recordLen).
		Order("search_time desc").
		Find(&records).Error; err != nil {
		wh.logger.LogError("fail to paginated weather record", map[string]interface{}{"package": "weather_history", "method": "GetPaginatedRecords", "user_id": userID, "error": err.Error()})
		return nil, apperror.ReturnStoreErr("WH01", err)
	}
	return records, nil
}

// DeleteRecord deletes a weather record by its ID.
func (wh *StoreGorm) DeleteRecord(id int) error {
	result := wh.db.Where("id = ?", id).Delete(&Record{})
	if result.Error != nil {
		wh.logger.LogError("fail to delete weather record", map[string]interface{}{"package": "weather_history", "method": "DeleteRecord", "record_id": id, "error": result.Error.Error()})
		return apperror.ReturnStoreErr("WH04", result.Error)
	}
	if result.RowsAffected == 0 {
		return apperror.WeatherHistoryRecordNotFound
	}
	return nil
}

// BulkDeleteRecord deletes multiple weather records.
func (wh *StoreGorm) BulkDeleteRecords(recordIDs []int) error {
	if result := wh.db.Where("id IN (?)", recordIDs).Delete(&Record{}); result.Error != nil {
		wh.logger.LogError("fail to bulk delete weather record", map[string]interface{}{"package": "weather_history", "method": "BulkDeleteRecords", "error": result.Error.Error()})
		return apperror.ReturnStoreErr("WH04", result.Error)
	} else if result.RowsAffected < int64(len(recordIDs)) {
		wh.logger.LogError("partial bul delete records", map[string]interface{}{"rec_ids": recordIDs})
		return nil
	}
	return nil
}
