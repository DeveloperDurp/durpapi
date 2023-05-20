package controller

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

type Controller struct {
	openaiClient *openai.Client
	unraidAPIKey string
	unraidURI    string

	config *configStruct
}

type configStruct struct {
	openaiApiKey string `json : "OPENAI_API_KEY"`
	unraidAPIKey string `json : "UNRAID_API_KEY"`
	unraidURI    string `json : "UNRAID_URI"`
}

func NewController() *Controller {
	err := godotenv.Load(".env")

	openaiApiKey := os.Getenv("OPENAI_API_KEY")
	openaiClient := openai.NewClient(openaiApiKey)
	unraidAPIKey := os.Getenv("UNRAID_API_KEY")
	unraidURI := os.Getenv("UNRAID_URI")

	if err != nil {
		fmt.Println(err.Error())
		//return err
	}
	return &Controller{
		openaiClient: openaiClient,
		unraidAPIKey: unraidAPIKey,
		unraidURI:    unraidURI,
	}
}

type Message struct {
	Message string `json:"message" example:"message"`
}
