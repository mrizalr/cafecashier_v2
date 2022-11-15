package usecase

import (
	"context"
	"encoding/json"

	"github.com/mrizalr/cafecashierpt2/domain"
	"github.com/mrizalr/cafecashierpt2/models"
	"github.com/mrizalr/cafecashierpt2/utils"
	"golang.org/x/crypto/bcrypt"
)

type ucaseAdmin struct {
	adminRepo domain.AdminRepository
}

func NewUcaseAdmin(adminRepo domain.AdminRepository) *ucaseAdmin {
	return &ucaseAdmin{adminRepo}
}

func (u *ucaseAdmin) Add(ctx context.Context, req *models.CreateNewAdminRequest) (domain.Admin, error) {
	result := new(domain.Admin)

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return *result, err
	}

	insertedID, err := u.adminRepo.Add(ctx, &domain.Admin{
		Username: req.Username,
		Password: string(hash),
		Role:     req.Role,
	})

	if err != nil {
		return *result, err
	}

	result.ID = int(insertedID)
	result.Username = req.Username
	result.Role = req.Role
	return *result, nil
}

func (u *ucaseAdmin) Login(ctx context.Context, req *models.AdminLoginRequest) (string, error) {
	admin, err := u.adminRepo.FindByUsername(context.Background(), req.Username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}

	result := utils.FormatToCreateNewAdminResponse(admin)
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	token := utils.Encode(jsonResult)

	return token, nil
}
