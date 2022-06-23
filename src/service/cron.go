package service

import (
	"log"
	"time"
)

var stopRecurringTasks chan bool = make(chan bool)

type recurringTask struct {
	Ticker  *time.Ticker
	Fn      func(*time.Ticker, chan bool)
	Stopper chan bool
}

func StartRecurringTasks() {
	log.Println("Starting recurring tasks")
	recurringTasks := make([]recurringTask, 0)

	recurringTasks = append(
		recurringTasks,
		recurringTask{
			Ticker:  time.NewTicker(2 * time.Second),
			Fn:      PortfolioRefresher,
			Stopper: make(chan bool),
		},
	)

	// Start recurring tasks
	for _, recurringTask := range recurringTasks {
		go recurringTask.Fn(recurringTask.Ticker, recurringTask.Stopper)
	}

	<-stopRecurringTasks

	// Stop recurring tasks once signal is reached
	for _, recurringTask := range recurringTasks {
		// Stop the ticker
		recurringTask.Ticker.Stop()

		// Stop the function execution
		recurringTask.Stopper <- true
	}
}

func StopRecurringTasks() {
	log.Println("Stopping recurring tasks")
	stopRecurringTasks <- true
}
