package domain

import (
	"context"

	"github.com/mrizalr/cafecashierpt2/models"
)

type Admin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role_id"`
}

type AdminRepository interface {
	Add(ctx context.Context, admin *Admin) (int64, error)
	FindByID(ctx context.Context, ID int) (Admin, error)
	FindByUsername(ctx context.Context, username string) (Admin, error)
}

type AdminUseCase interface {
	Add(ctx context.Context, req *models.CreateNewAdminRequest) (Admin, error)
	Login(ctx context.Context, req *models.AdminLoginRequest) (string, error)
}
