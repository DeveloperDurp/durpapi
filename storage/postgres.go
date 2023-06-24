package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gitlab.com/DeveloperDurp/DurpAPI/model"
)

func Connect(config model.DBConfig) (*model.Repository, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = runMigrations(db)
	if err != nil {
		return nil, err
	}
	return &model.Repository{
		DB: db,
	}, nil
}

func runMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&model.DadJoke{})
	if err != nil {
		return err
	}
	return nil
}
