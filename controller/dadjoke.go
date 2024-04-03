package controller

import (
	"encoding/json"
	"net/http"

	"gitlab.com/DeveloperDurp/DurpAPI/model"
	"gitlab.com/DeveloperDurp/DurpAPI/service"
)

// GetDadJoke godoc
//
//	@Summary		Get dadjoke
//	@Description	get a dad joke
//	@Tags			DadJoke
//	@Accept			json
//	@Produce		application/json
//	@Success		200	{object}	model.Message	"response"
//	@failure		500	{object}	model.Message	"error"
//	@Router			/jokes/dadjoke [get]
func (c *Controller) GetDadJoke(w http.ResponseWriter, r *http.Request) {
	joke, err := service.GetRandomDadJoke(c.Db.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	message := model.Message{
		Message: joke,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

// PostDadJoke godoc
//
//	@Summary		Generate dadjoke
//	@Description	create a dad joke
//	@Tags			DadJoke
//	@Accept			json
//	@Produce		application/json
//	@Param			joke	query		string			true	"Dad Joke you wish to enter into database"
//	@Success		200		{object}	model.Message	"response"
//	@failure		500		{object}	model.Message	"error"
//	@Router			/jokes/dadjoke [post]
func (c *Controller) PostDadJoke(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	var req model.DadJoke

	if contentType == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}
	} else {
		queryParams := r.URL.Query()
		req.JOKE = queryParams.Get("joke")
	}

	err := service.PostDadJoke(c.Db.DB, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	message := model.Message{
		Message: "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

// DeleteDadJoke godoc
//
//	@Summary		Generate dadjoke
//	@Description	create a dad joke
//	@Tags			DadJoke
//	@Accept			json
//	@Produce		application/json
//	@Param			joke	query		string			true	"Dad joke you wish to delete from the database"
//	@Success		200		{object}	model.Message	"response"
//	@failure		500		{object}	model.Message	"error"
//	@Router			/jokes/dadjoke [delete]
func (c *Controller) DeleteDadJoke(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	var req model.DadJoke

	if contentType == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}
	} else {
		queryParams := r.URL.Query()
		req.JOKE = queryParams.Get("joke")
	}

	err := service.DeleteDadJoke(c.Db.DB, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	message := model.Message{
		Message: "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
