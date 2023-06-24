package service

import (
	"errors"
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

func PostDadJoke(db *gorm.DB, joke model.DadJoke) error {
	jokes, err := getDadJokes(db)
	if err != nil {
		return err
	}

	found := false
	for _, i := range jokes {
		if i.JOKE == joke.JOKE {
			found = true
			break
		}
	}

	if found {
		return errors.New("Joke is already in database")
	} else {
		err = db.Create(&joke).Error
		if err != nil {
			return err
		}
		return nil
	}
}

func DeleteDadJoke(db *gorm.DB, joke model.DadJoke) error {
	check := &model.DadJoke{}
	db.Where("joke = ?", joke.JOKE).First(check)
	if check.JOKE == "" {
		return errors.New("Joke does not exist")
	}

	err := db.Where("joke = ?", joke.JOKE).Delete(joke).Error
	if err != nil {
		return err
	}
	return nil
}

func getDadJokes(db *gorm.DB) ([]model.DadJoke, error) {
	req := []model.DadJoke{}

	err := db.Find(&req).Error

	return req, err
}
