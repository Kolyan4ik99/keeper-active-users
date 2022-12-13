package model

import (
	"net"
	"time"
)

type User struct {
	IpAddr     net.IP    `json:"ip_address,omitempty"`
	Since      int64     `json:"since,omitempty"`
	ExpiryTime time.Time `json:"-"`
}

func NewUser(ipAddr net.IP, expiryTime time.Duration) User {
	now := time.Now()
	return User{
		IpAddr:     ipAddr,
		Since:      now.Unix(),
		ExpiryTime: now.Add(expiryTime),
	}
}
