package domain

import (
	"context"

	"github.com/mrizalr/cafecashierpt2/models"
)

type Admin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminRepository interface {
	Add(ctx context.Context, admin *Admin) (int64, error)
	FindByID(ctx context.Context, ID int) (Admin, error)
}

type AdminUseCase interface {
	Add(ctx context.Context, req *models.CreateNewAdminRequest) (Admin, error)
}
