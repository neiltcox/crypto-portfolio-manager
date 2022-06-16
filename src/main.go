package main

import (
	"log"

	"github.com/neiltcox/coinbake/config"
	"github.com/neiltcox/coinbake/database"
	"github.com/neiltcox/coinbake/service"
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

	sandbox(config)
}

func sandbox(cfg config.Config) {
	log.Println("Sandbox starting")

	user := service.FindUserByEmailAddress("neiltcox@gmail.com")

	//log.Printf("%#v", user)

	//service.RefreshMarketData(cfg)

	exchangeConnections := service.FindPortfoliosByUserId(user.ID)
	for _, exchangeConnection := range exchangeConnections {
		log.Printf("exchange conection: %d", exchangeConnection.ID)
		strategy := service.FindStrategyByExchangeConnectionId(exchangeConnection.ID)
		if strategy == nil {
			log.Printf("strategy is nil")
			continue
		}

		schedule, err := strategy.GenerateSchedule(&exchangeConnection)
		if err != nil {
			log.Printf("Could not generate schedule: %s", err)
		}

		for _, scheduleItem := range schedule.Items {
			log.Printf("%s -> %f", scheduleItem.Asset.Symbol, scheduleItem.Amount)
		}
	}

	log.Println("Sandbox concluded")
}
