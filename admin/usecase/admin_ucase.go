package usecase

import (
	"context"

	"github.com/mrizalr/cafecashierpt2/domain"
	"github.com/mrizalr/cafecashierpt2/models"
	"golang.org/x/crypto/bcrypt"
)

type ucaseAdmin struct {
	adminRepo domain.AdminRepository
}

func NewUcaseAdmin(adminRepo domain.AdminRepository) *ucaseAdmin {
	return &ucaseAdmin{adminRepo}
}

func (u *ucaseAdmin) Add(ctx context.Context, req *models.CreateNewAdminRequest) (domain.Admin, error) {
	result := domain.Admin{}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return result, err
	}

	insertedID, err := u.adminRepo.Add(ctx, &domain.Admin{
		Username: req.Username,
		Password: string(hash),
		Role:     req.Role,
	})

	if err != nil {
		return result, err
	}

	result.ID = int(insertedID)
	result.Username = req.Username
	result.Role = req.Role
	return result, nil
}
