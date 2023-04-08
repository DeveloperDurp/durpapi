package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

var (
	// Declare global variables
	apiKey string

	config *configStruct
)

type configStruct struct {
	apiKey string `json : "API_KEY"`
}

func init() {
	// Load values for global variables from environment variables
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(err.Error())
		//return err
	}
	apiKey = os.Getenv("API_KEY")

}

// GeneralOpenAI godoc
//
//	@Summary		Gerneral ChatGPT
//	@Description	Ask ChatGPT a general question
//	@Tags			openai
//	@Accept			json
//	@Produce		plain
//	@Param			message	query		string		true	"Ask ChatGPT a general question"
//	@Success		200	{string}	string	"response"
//	@Router			/openai/general [get]
func (c *Controller) GeneralOpenAI(ctx *gin.Context) {
	message := ctx.Query("message")

	result, err := createChatCompletion(message)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.String(http.StatusOK, result)
}

// TravelAgentOpenAI godoc
//
//	@Summary		Travel Agent ChatGPT
//	@Description	Ask ChatGPT for suggestions as if it was a travel agent
//	@Tags			openai
//	@Accept			json
//	@Produce		plain
//	@Param			message	query		string		true	"Ask ChatGPT for suggestions as a travel agent"
//	@Success		200	{string}	string	"response"
//	@Router			/openai/travelagent [get]
func (c *Controller) TravelAgentOpenAI(ctx *gin.Context) {
	message := "want you to act as a travel guide. I will give you my location and you will give me suggestions " + ctx.Query("message")

	result, err := createChatCompletion(message)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.String(http.StatusOK, result)
}

func createChatCompletion(message string) (string, error) {

	var client = openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
