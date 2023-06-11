package controller

import (
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type Controller struct {
	openaiClient *openai.Client
	unraidAPIKey string
	unraidURI    string
	ClientID     string
	ClientSecret string
	RedirectURL  string
	AuthURL      string
	TokenURL     string
}

func NewController() *Controller {

	openaiApiKey := os.Getenv("OPENAI_API_KEY")
	unraidAPIKey := os.Getenv("UNRAID_API_KEY")
	unraidURI := os.Getenv("UNRAID_URI")
	ClientID := os.Getenv("ClientID")
	ClientSecret := os.Getenv("ClientSecret")
	RedirectURL := os.Getenv("RedirectURL")
	AuthURL := os.Getenv("AuthURL")
	TokenURL := os.Getenv("TokenURL")

	return &Controller{
		openaiClient: openai.NewClient(openaiApiKey),
		unraidAPIKey: unraidAPIKey,
		unraidURI:    unraidURI,
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		RedirectURL:  RedirectURL,
		AuthURL:      AuthURL,
		TokenURL:     TokenURL,
	}
}
