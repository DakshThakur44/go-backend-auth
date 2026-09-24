package db

import (
	"auth-backend/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(dbPath string) (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, err
	}

	err = database.AutoMigrate(
		&models.User{},
		&models.Session{},
		&models.RefreshToken{},
	)
	return database, err
}
