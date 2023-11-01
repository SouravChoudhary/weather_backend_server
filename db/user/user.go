package user

import (
	"time"
)

type User struct {
	ID           int       `gorm:"primaryKey"`
	Username     string    `gorm:"not null;unique"`
	PasswordHash string    `gorm:"not null"`
	DateOfBirth  time.Time `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
