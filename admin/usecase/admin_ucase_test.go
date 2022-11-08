package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/mrizalr/cafecashierpt2/domain"
	"github.com/mrizalr/cafecashierpt2/domain/mocks"
	"github.com/mrizalr/cafecashierpt2/models"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("Test success add new admin", func(t *testing.T) {
		mock := new(mocks.MysqlAdminRepository)
		ucaseAdmin := ucaseAdmin{mock}

		mock.On("Add", context.Background(), &domain.Admin{
			Username: "test",
			Password: "test",
			Role:     "test",
		}).Return(1, nil)

		mock.On("FindByID", context.Background(), 1).Return(domain.Admin{
			ID:       1,
			Username: "test",
			Role:     "test",
		}, nil)

		result, err := ucaseAdmin.Add(context.Background(), &models.CreateNewAdminRequest{
			Username: "test",
			Password: "test",
			Role:     "test",
		})

		mock.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, domain.Admin{
			ID:       1,
			Username: "test",
			Password: "",
			Role:     "test",
		}, result)
	})

	t.Run("Test failed add new admin", func(t *testing.T) {
		mock := new(mocks.MysqlAdminRepository)
		ucaseAdmin := ucaseAdmin{mock}

		mock.On("Add", context.Background(), &domain.Admin{
			Username: "test",
			Password: "test",
			Role:     "test",
		}).Return(0, errors.New("Error inserting data"))

		result, err := ucaseAdmin.Add(context.Background(), &models.CreateNewAdminRequest{
			Username: "test",
			Password: "test",
			Role:     "test",
		})

		mock.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Equal(t, domain.Admin{}, result)
		t.Log(err)
	})
}
