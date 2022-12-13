package service

import (
	"context"
	"net"
	"time"

	"keeper-active-users/internal/model"
)

type UserRepository interface {
	GetByIP(ip net.IP) (model.User, bool)
	Save(ip net.IP, expiryTime time.Duration)
	UpdateByIP(ip net.IP, expiryTime time.Duration)
}

type User struct {
	expiryTime time.Duration
	repo       UserRepository
}

func NewUser(expiryTime time.Duration, repo UserRepository) *User {
	return &User{expiryTime: expiryTime, repo: repo}
}

func (u *User) Ping(_ context.Context, ip net.IP) error {
	_, exist := u.repo.GetByIP(ip)
	if !exist {
		u.repo.Save(ip, u.expiryTime)
		return nil
	}
	u.repo.UpdateByIP(ip, u.expiryTime)
	return nil
}
