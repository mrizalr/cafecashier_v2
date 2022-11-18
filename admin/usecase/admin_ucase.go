package usecase

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/mrizalr/cafecashierpt2/domain"
	"github.com/mrizalr/cafecashierpt2/models"
	"github.com/mrizalr/cafecashierpt2/utils"
	"golang.org/x/crypto/bcrypt"
)

type ucaseAdmin struct {
	adminRepo domain.AdminRepository
}

func NewUcaseAdmin(adminRepo domain.AdminRepository) *ucaseAdmin {
	usecase := &ucaseAdmin{adminRepo}
	if _, err := adminRepo.FindByUsername(context.Background(), "owner"); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			usecase.Add(context.Background(), &models.CreateNewAdminRequest{
				Username: "owner",
				Password: "owner123",
				Role:     1,
			})
		}
	}

	return usecase
}

func (u *ucaseAdmin) Add(ctx context.Context, req *models.CreateNewAdminRequest) (models.Admin, error) {
	result := new(models.Admin)

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

	role, err := u.adminRepo.FindAdminRoleByID(ctx, req.Role)
	if err != nil {
		return *result, err
	}

	result.ID = int(insertedID)
	result.Username = req.Username
	result.Role = role

	return *result, nil
}

func (u *ucaseAdmin) Login(ctx context.Context, req *models.AdminLoginRequest) (string, error) {
	admin, err := u.adminRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}

	response := utils.FormatToCreateNewAdminResponse(admin)
	jsonResult, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	token := utils.Encode(jsonResult)

	return token, nil
}

func (u *ucaseAdmin) GetAdmins(ctx context.Context) ([]models.Admin, error) {
	return u.adminRepo.FindAll(ctx)
}
