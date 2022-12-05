package service

import (
	"context"

	"keeper-active-users/internal/model"
)

type AdminRepository interface {
	GetUsers() []*model.User
}

type Admin struct {
	repo AdminRepository
}

func NewAdmin(repo AdminRepository) *Admin {
	return &Admin{repo: repo}
}

func (a *Admin) GetUsers(_ context.Context) ([]*model.User, error) {
	return a.repo.GetUsers(), nil
}
