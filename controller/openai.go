package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gitlab.com/DeveloperDurp/DurpAPI/model"
)

type ChatRequest struct {
	Message string `json:"message"`
}

// Response struct to unmarshal the JSON response
type Response struct {
	Response string `json:"response"`
}

// GeneralOpenAI godoc
//
//	@Summary		Gerneral ChatGPT
//	@Description	Ask ChatGPT a general question
//	@Tags			openai
//	@Accept			json
//	@Produce		application/json
//	@Param			message	query		string			true	"Ask ChatGPT a general question"
//	@Success		200		{object}	model.Message	"response"
//
//	@failure		400		{object}	model.Message	"error"
//
//	@Router			/openai/general [get]
func (c *Controller) GeneralOpenAI(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	var req ChatRequest

	if contentType == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}
	} else {
		queryParams := r.URL.Query()
		req.Message = queryParams.Get("message")
	}

	result, err := c.createChatCompletion(req.Message, "mistral:instruct")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	message := model.Message{
		Message: result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

// TravelAgentOpenAI godoc
//
//	@Summary		Travel Agent ChatGPT
//	@Description	Ask ChatGPT for suggestions as if it was a travel agent
//	@Tags			openai
//	@Accept			json
//	@Produce		application/json
//	@Param			message	query		string			true	"Ask ChatGPT for suggestions as a travel agent"
//	@Success		200		{object}	model.Message	"response"
//	@failure		400		{object}	model.Message	"error"
//	@Router			/openai/travelagent [get]
func (c *Controller) TravelAgentOpenAI(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	var req ChatRequest

	if contentType == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}
	} else {
		queryParams := r.URL.Query()
		req.Message = queryParams.Get("message")
	}

	req.Message = "I want you to act as a travel guide. I will give you my location and you will give me suggestions. " + req.Message

	result, err := c.createChatCompletion(req.Message, "mistral:instruct")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	message := model.Message{
		Message: result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func (c *Controller) createChatCompletion(message string, model string) (string, error) {
	// Define the request body
	requestBody := map[string]interface{}{
		"model":  model,
		"prompt": message,
		"stream": false,
	}

	// Convert the request body to JSON
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error encoding request body: %v", err)
	}

	// Send a POST request to the specified URL with the request body
	response, err := http.Post(
		"http://"+c.Cfg.LlamaURL+"/api/generate",
		"application/json",
		bytes.NewBuffer(requestBodyBytes),
	)
	if err != nil {
		return "", fmt.Errorf("error sending POST request: %v", err)
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	// Unmarshal the JSON response
	var resp Response
	if err := json.Unmarshal(responseBody, &resp); err != nil {
		return "", fmt.Errorf("error decoding response body: %v", err)
	}

	// Return the response
	return resp.Response, nil
}
