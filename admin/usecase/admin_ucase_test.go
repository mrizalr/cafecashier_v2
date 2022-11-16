package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/mrizalr/cafecashierpt2/domain"
	"github.com/mrizalr/cafecashierpt2/domain/mocks"
	"github.com/mrizalr/cafecashierpt2/models"
	"github.com/mrizalr/cafecashierpt2/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestAdd(t *testing.T) {
	fields := struct {
		username string
		password string
		role     int
	}{
		username: "test",
		password: "test123",
		role:     1,
	}

	t.Run("Test success add new admin", func(t *testing.T) {
		mockRepo := new(mocks.MysqlAdminRepository)
		ucaseAdmin := ucaseAdmin{mockRepo}

		mockRepo.On("Add", context.Background(), mock.AnythingOfType("*domain.Admin")).Return(1, nil)

		result, err := ucaseAdmin.Add(context.Background(), &models.CreateNewAdminRequest{
			Username: fields.username,
			Password: fields.password,
			Role:     fields.role,
		})

		mockRepo.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, domain.Admin{
			ID:       1,
			Username: fields.username,
			Password: "",
			Role:     fields.role,
		}, result)
	})

	t.Run("Test failed add new admin", func(t *testing.T) {
		mockRepo := new(mocks.MysqlAdminRepository)
		ucaseAdmin := ucaseAdmin{mockRepo}

		mockRepo.On("Add", context.Background(), mock.AnythingOfType("*domain.Admin")).Return(0, errors.New("Error inserting data"))

		result, err := ucaseAdmin.Add(context.Background(), &models.CreateNewAdminRequest{
			Username: fields.username,
			Password: fields.password,
			Role:     fields.role,
		})

		mockRepo.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Equal(t, domain.Admin{}, result)
	})
}

func Test_ucaseAdmin_Login(t *testing.T) {
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

	mockRepo.On("FindByUsername", context.Background(), mock.AnythingOfType("string")).Return(domain.Admin{
		ID:       1,
		Username: "admin",
		Password: string(h),
		Role:     1,
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

	res := domain.Admin{}
	j, err := utils.Decode(token)
	assert.NoError(t, err)
	err = json.Unmarshal(j, &res)
	assert.NoError(t, err)

	assert.Equal(t, admin.ID, res.ID)
	assert.Equal(t, admin.Username, res.Username)
	assert.Equal(t, admin.Role, res.Role)
}
