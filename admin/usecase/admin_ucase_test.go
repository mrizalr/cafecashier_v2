package usecase

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/mrizalr/cafecashierpt2/domain"
	"github.com/mrizalr/cafecashierpt2/domain/mocks"
	"github.com/mrizalr/cafecashierpt2/models"
	"github.com/mrizalr/cafecashierpt2/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func Test_UcaseAdmin_Add(t *testing.T) {
	mockRepo := new(mocks.MysqlAdminRepository)
	ucaseAdmin := ucaseAdmin{mockRepo}

	mockRepo.On("Add", context.Background(), mock.AnythingOfType("*domain.Admin")).Return(1, nil)
	mockRepo.On("FindAdminRoleByID", context.Background(), 1).Return("super admin", nil)

	result, err := ucaseAdmin.Add(context.Background(), &models.CreateNewAdminRequest{
		Username: "admin",
		Password: "admin123",
		Role:     1,
	})

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, models.Admin{
		ID:       1,
		Username: "admin",
		Role:     "super admin",
	}, result)
}

func Test_UcaseAdmin_Login(t *testing.T) {
	mockRepo := new(mocks.MysqlAdminRepository)
	usecaseAdmin := ucaseAdmin{
		adminRepo: mockRepo,
	}

	admin := domain.Admin{
		ID:       1,
		Username: "admin",
		Password: "admin123",
		Role:     1,
	}

	h, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	assert.NoError(t, err)

	mockRepo.On("FindByUsername", context.Background(), mock.AnythingOfType("string")).Return(models.Admin{
		ID:       1,
		Username: "admin",
		Password: string(h),
		Role:     "super admin",
	}, nil)

	token, err := usecaseAdmin.Login(context.Background(), &models.AdminLoginRequest{
		Username: "admin",
		Password: "admin123",
	})
	assert.NoError(t, err)

	err = bcrypt.CompareHashAndPassword(h, []byte(admin.Password))
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
	assert.NotEqual(t, "", token)

	res := models.Admin{}
	j, err := utils.Decode(token)
	assert.NoError(t, err)
	err = json.Unmarshal(j, &res)
	assert.NoError(t, err)

	assert.Equal(t, admin.ID, res.ID)
	assert.Equal(t, admin.Username, res.Username)
	assert.Equal(t, "super admin", res.Role)
}

func Test_UcaseAdmin_GetAdmins(t *testing.T) {
	mockRepo := new(mocks.MysqlAdminRepository)
	adminUcase := ucaseAdmin{
		adminRepo: mockRepo,
	}

	mockRepo.On("FindAll", context.Background()).Return([]models.Admin{{
		ID:       1,
		Username: "admin",
		Password: "",
		Role:     "super admin",
	}, {
		ID:       2,
		Username: "finance",
		Password: "",
		Role:     "finance",
	}}, nil)

	admins, err := adminUcase.GetAdmins(context.Background())
	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Len(t, admins, 2)
}
