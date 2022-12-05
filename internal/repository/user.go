package repository

import (
	"net"
	"sync"
	"time"

	"keeper-active-users/internal/model"
)

type User struct {
	users map[string]*model.User // [ip_address : user]
	mu    sync.RWMutex
}

func NewUser() *User {
	return &User{users: make(map[string]*model.User)}
}

func (u *User) GetByIP(ip net.IP) (*model.User, bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	user, exist := u.users[ip.String()]
	return user, exist
}

func (u *User) Save(ip net.IP, expiryTime time.Duration) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.users[ip.String()] = model.NewUser(ip, expiryTime)
}

func (u *User) UpdateByIP(ip net.IP, expiryTime time.Duration) {
	u.mu.Lock()
	defer u.mu.Unlock()
	user := u.users[ip.String()]
	user.ExpiryTime = time.Now().Add(expiryTime)
}

func (u *User) GetUsers() []*model.User {
	u.mu.RLock()
	defer u.mu.RUnlock()
	var users []*model.User
	for _, user := range u.users {
		users = append(users, user)
	}
	return users
}

func (u *User) DeleteByIP(ip net.IP) {
	u.mu.Lock()
	defer u.mu.Unlock()
	delete(u.users, ip.String())
}
