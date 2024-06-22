package dadjoke

import (
	"errors"
	"math/rand"
)

func (h *Handler) GetRandomDadJoke() (string, error) {
	jokes, err := h.getDadJokes()
	if err != nil {
		return "", err
	}

	randomIndex := rand.Intn(len(jokes))

	randomElement := jokes[randomIndex]

	return randomElement.JOKE, err
}

func (h *Handler) PostDadJoke(joke DadJoke) error {
	jokes, err := h.getDadJokes()
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
		err = h.db.Create(&joke).Error
		if err != nil {
			return err
		}
		return nil
	}
}

func (h *Handler) DeleteDadJoke(joke DadJoke) error {
	check := &DadJoke{}
	h.db.Where("joke = ?", joke.JOKE).First(check)
	if check.JOKE == "" {
		return errors.New("Joke does not exist")
	}

	err := h.db.Where("joke = ?", joke.JOKE).Delete(joke).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) getDadJokes() ([]DadJoke, error) {
	req := []DadJoke{}

	err := h.db.Find(&req).Error

	return req, err
}
