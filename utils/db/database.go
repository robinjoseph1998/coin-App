package utils

import (
	"coin-App/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=123456 dbname=coindb port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db

	err = DB.AutoMigrate(
		&models.Coin{},
		&models.ExpiredCoinLogs{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
