package service

import (
	"log"

	"github.com/jaksonkallio/coinbake/database"
)

func InitModels() {
	log.Println("Initializing models in database")

	database.Handle().AutoMigrate(
		&User{},
		&Strategy{},
	)
}
