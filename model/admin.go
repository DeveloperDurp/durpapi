package model

import (
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

type Config struct {
	OpenaiClient openai.Client
	OpenaiApiKey string `env:"openai_api_key"`
	UnraidAPIKey string `env:"unraid_api_key"`
	UnraidURI    string `env:"unraid_uri"`
	Host         string `env:"host"`
	Version      string `env:"version"`
	Groupsenv    string `env:"groupsenv"`
	JwksURL      string `env:"jwksurl"`
}

type DBConfig struct {
	Host     string `env:"db_host"`
	Port     string `env:"db_port"`
	Password string `env:"db_pass"`
	User     string `env:"db_user"`
	DBName   string `env:"db_name"`
	SSLMode  string `env:"db_sslmode"`
}

type Repository struct {
	DB *gorm.DB
}
