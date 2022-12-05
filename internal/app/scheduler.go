package app

import (
	"log"
	"time"

	"keeper-active-users/internal/config"
	"keeper-active-users/internal/repository"
)

type Scheduler struct {
	cfg config.Scheduler
}

func NewScheduler(cfg config.Scheduler) *Scheduler {
	return &Scheduler{cfg: cfg}
}

func (s *Scheduler) Start(userRepository *repository.User) {
	ticker := time.NewTicker(s.cfg.TickTime)
	for {
		select {
		case <-ticker.C:
			log.Println("Scheduler is running")
			users := userRepository.GetUsers()

			for _, user := range users {
				if time.Now().Sub(user.ExpiryTime) >= 0 {

					userRepository.DeleteByIP(user.IpAddr)
					log.Printf("User with ip=[%s] is deleted\n", user.IpAddr)
				}
			}
			log.Println("Scheduler is down")
		}
	}
}
