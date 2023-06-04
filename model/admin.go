package model

import "os"

var (
	OpenaiApiKey = os.Getenv("OPENAI_API_KEY")
	UnraidAPIKey = os.Getenv("UNRAID_API_KEY")
	UnraidURI    = os.Getenv("UNRAID_URI")
	ClientID     = os.Getenv("ClientID")
	ClientSecret = os.Getenv("ClientSecret")
	RedirectURL  = os.Getenv("RedirectURL")
	AuthURL      = os.Getenv("AuthURL")
	TokenURL     = os.Getenv("TokenURL")
	Host         = os.Getenv("host")
	Version      = os.Getenv("version")
)
