package weather_history

import (
	"time"
)

type Record struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	UserID      int       `gorm:"not null" json:"user_id"`
	CityName    string    `gorm:"not null" json:"city_name"`
	WeatherData string    `gorm:"not null" json:"weather_data"`
	SearchTime  time.Time `gorm:"autoCreateTime" json:"search_time"`
}
