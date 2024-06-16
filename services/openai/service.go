package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h *Handler) createChatCompletion(message string, model string) (string, error) {
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
		"http://"+h.LlamaURL+"/api/generate",
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
