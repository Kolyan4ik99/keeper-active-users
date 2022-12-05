package main

import (
	"log"
	"sync"
	"time"

	"keeper-active-users/internal/app"
	"keeper-active-users/internal/config"
	"keeper-active-users/internal/repository"
)

func main() {
	userRepository := repository.NewUser()

	expiryTime := time.Minute * 30

	wg := sync.WaitGroup{}

	server := app.NewServer(config.Server{
		Addr:           "127.0.0.1:8080",
		UserExpiryTime: expiryTime,
	})

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.Start(userRepository)
		if err != nil {
			log.Fatal(err)
		}
	}()

	scheduler := app.NewScheduler(config.Scheduler{
		TickTime: time.Minute,
	})

	wg.Add(1)
	go func() {
		defer wg.Done()
		scheduler.Start(userRepository)
	}()

	wg.Wait()

}
