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
	FindByID(ctx context.Context, ID int) (models.Admin, error)
	FindByUsername(ctx context.Context, username string) (models.Admin, error)
	FindAll(ctx context.Context) ([]models.Admin, error)
	FindAdminRoleByID(ctx context.Context, ID int) (string, error)
}

type AdminUseCase interface {
	Add(ctx context.Context, req *models.CreateNewAdminRequest) (models.Admin, error)
	Login(ctx context.Context, req *models.AdminLoginRequest) (string, error)
	GetAdmins(ctx context.Context) ([]models.Admin, error)
}
