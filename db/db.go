package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
dsn := "host=localhost user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable"
*/
func ConnectToDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
