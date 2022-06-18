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

	service.Serve()

	log.Println("Fully stopped.")
}

// TODO: remove all of this junk
func sandbox(cfg config.Config) {
	log.Println("Sandbox starting")

	user := service.FindUserByEmailAddress("jaksonkallio@gmail.com")

	//log.Printf("%#v", user)

	//service.RefreshMarketData(cfg)

	portfolios := service.FindPortfoliosByUserId(user.ID)
	for _, portfolio := range portfolios {
		log.Printf("exchange conection: %d", portfolio.ID)
		strategy := service.FindStrategyByPortfolioId(portfolio.ID)
		if strategy == nil {
			log.Printf("strategy is nil")
			continue
		}

		rebalanceMovements, err := strategy.RebalanceMovements(&portfolio)
		if err != nil {
			log.Printf("Could not generate rebalance movements: %s", err)
		}

		for _, rebalanceMovement := range rebalanceMovements.Movements {
			log.Printf(
				"%s tgwt: %f vldf: %f atdf: %f",
				rebalanceMovement.Asset.Symbol,
				rebalanceMovement.WeightProportion,
				rebalanceMovement.ValuationDiff,
				rebalanceMovement.BalanceDiff(),
			)
		}
	}

	log.Println("Sandbox concluded")
}
