package config

import "time"

type Server struct {
	Addr           string
	UserExpiryTime time.Duration
}
