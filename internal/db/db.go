package db

import (
	"github.com/anazibinurasheed/dmart-auth-svc/internal/config"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(cfg config.Config) (*gorm.DB, error) {
	dsn := cfg.DbURL

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true, Logger: logger.Default.LogMode(logger.Info)})
	db.Debug()

	if err != nil {
		return &gorm.DB{}, err
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Address{},
	)
	if err != nil {
		return &gorm.DB{}, err
	}

	return db, nil
}
