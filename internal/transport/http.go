package transport

import (
	"net/http"

	"github.com/gorilla/mux"
)

type AdminHandler interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
}

type UserHandler interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	admin AdminHandler
	user  UserHandler
}

func NewHandler(admin AdminHandler, user UserHandler) *Handler {
	return &Handler{admin: admin, user: user}
}

func (h *Handler) GetRoute() *mux.Router {
	route := mux.NewRouter()

	admin := route.PathPrefix("/admin").Subrouter()
	{
		admin.HandleFunc("/users", h.admin.GetUsers).Methods(http.MethodGet)
	}

	user := route.PathPrefix("/user").Subrouter()
	{
		user.HandleFunc("/ping", h.user.Ping).Methods(http.MethodGet)
	}

	return route
}
