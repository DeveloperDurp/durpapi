package service

import (
	"math/rand"

	"gorm.io/gorm"

	"gitlab.com/DeveloperDurp/DurpAPI/model"
)

func GetRandomDadJoke(db *gorm.DB) (string, error) {
	jokes, err := getDadJokes(db)
	if err != nil {
		return "", err
	}

	randomIndex := rand.Intn(len(jokes))

	randomElement := jokes[randomIndex]

	return randomElement.JOKE, err
}

func getDadJokes(db *gorm.DB) ([]model.DadJoke, error) {
	req := []model.DadJoke{}

	err := db.Find(&req).Error

	return req, err
}
