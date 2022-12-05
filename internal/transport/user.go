package transport

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"
)

type UserService interface {
	Ping(ctx context.Context, ip net.IP) error
}

type User struct {
	userService UserService
}

func NewUser(userService UserService) *User {
	return &User{userService: userService}
}

func (u *User) Ping(w http.ResponseWriter, r *http.Request) {
	ipAndPort := strings.Split(r.RemoteAddr, ":")

	ip := net.ParseIP(ipAndPort[0])

	err := u.userService.Ping(r.Context(), ip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Update ip=[%s]\n", ip)
	w.WriteHeader(http.StatusOK)
}
