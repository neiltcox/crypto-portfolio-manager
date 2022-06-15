package main

import (
	"log"

	"github.com/jaksonkallio/coinbake/config"
	"github.com/jaksonkallio/coinbake/database"
	"github.com/jaksonkallio/coinbake/service"
)

func main() {
	config, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("could not load config: %s", err)
	}

	err = database.Connect(config)
	if err != nil {
		log.Fatalf("could not connect to database: %s", err)
	}

	service.InitModels()

	log.Println("success")
}
