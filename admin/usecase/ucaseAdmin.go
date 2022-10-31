package usecase

import (
	"context"

	"github.com/mrizalr/cafecashierpt2/domain"
	"github.com/mrizalr/cafecashierpt2/models"
)

type ucaseAdmin struct {
	adminRepo domain.AdminRepository
}

func NewUcaseAdmin(adminRepo domain.AdminRepository) *ucaseAdmin {
	return &ucaseAdmin{adminRepo}
}

func (u *ucaseAdmin) Add(ctx context.Context, req *models.CreateNewAdminRequest) (res domain.Admin, err error) {
	newAdmin := domain.Admin{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}

	insertedID, err := u.adminRepo.Add(ctx, &newAdmin)
	if err != nil {
		return
	}

	res, err = u.adminRepo.FindByID(ctx, int(insertedID))
	return
}
