package model

import (
	"gorm.io/gorm"
)

type Config struct {
	Host      string `env:"host"`
	Version   string `env:"version"`
	Groupsenv string `env:"groupsenv"`
	JwksURL   string `env:"jwksurl"`
	LlamaURL  string `env:"llamaurl"`
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
