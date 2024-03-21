package controller

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"

	"gitlab.com/DeveloperDurp/DurpAPI/model"
	"gitlab.com/DeveloperDurp/DurpAPI/storage"
)

type Controller struct {
	Cfg   model.Config
	Dbcfg model.DBConfig
	Db    model.Repository
}

func NewController() *Controller {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("no env file found")
	}

	controller := &Controller{
		Cfg:   model.Config{},
		Dbcfg: model.DBConfig{},
	}

	err = env.Parse(&controller.Cfg)
	if err != nil {
		log.Fatalf("unable to parse environment variables: %e", err)
	}
	err = env.Parse(&controller.Dbcfg)
	if err != nil {
		log.Fatalf("unable to parse database variables: %e", err)
	}

	Db, err := storage.Connect(controller.Dbcfg)
	if err != nil {
		panic("Failed to connect to database")
	}
	controller.Db = *Db

	return controller
}
