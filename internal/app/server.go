package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"keeper-active-users/internal/config"
	"keeper-active-users/internal/repository"
	"keeper-active-users/internal/service"
	"keeper-active-users/internal/transport"
)

type Server struct {
	cfg config.Server
}

func NewServer(cfg config.Server) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Start(userRepository *repository.User) error {
	log.Println("Server initialization")

	adminService := service.NewAdmin(userRepository)
	userService := service.NewUser(s.cfg.UserExpiryTime, userRepository)

	adminTransport := transport.NewAdmin(adminService)
	userTransport := transport.NewUser(userService)

	handler := transport.NewHandler(adminTransport, userTransport)

	srv := &http.Server{
		Handler:      handler.GetRoute(),
		Addr:         s.cfg.Addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	doneSignal := make(chan os.Signal, 1)
	signal.Notify(doneSignal, os.Interrupt)

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("Server is started")

	<-doneSignal
	return fmt.Errorf("catch interrupt signal")
}
