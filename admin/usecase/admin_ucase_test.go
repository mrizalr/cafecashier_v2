package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/mrizalr/cafecashierpt2/domain"
	"github.com/mrizalr/cafecashierpt2/domain/mocks"
	"github.com/mrizalr/cafecashierpt2/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
