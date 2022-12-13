package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"keeper-active-users/internal/model"
)

type AdminService interface {
	GetUsers(ctx context.Context) ([]model.User, error)
}

type Admin struct {
	adminService AdminService
}

func NewAdmin(adminService AdminService) *Admin {
	return &Admin{adminService: adminService}
}

func (a *Admin) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := a.adminService.GetUsers(r.Context())
	if err != nil {
		w.Write([]byte("{}"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		w.Write([]byte("{}"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	msgBytes, err := json.Marshal(users)
	if err != nil {
		w.Write([]byte("{}"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(msgBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
