package mocks

import (
	"context"

	"github.com/mrizalr/cafecashierpt2/domain"
	"github.com/mrizalr/cafecashierpt2/models"
	"github.com/stretchr/testify/mock"
)

type MysqlAdminRepository struct {
	mock.Mock
}

func (r MysqlAdminRepository) Add(ctx context.Context, admin *domain.Admin) (int64, error) {
	arguments := r.Called(ctx, admin)
	return int64(arguments.Int(0)), arguments.Error(1)
}

func (r MysqlAdminRepository) FindByID(ctx context.Context, id int) (models.Admin, error) {
	arguments := r.Called(ctx, id)
	return arguments.Get(0).(models.Admin), arguments.Error(1)
}

func (r MysqlAdminRepository) FindByUsername(ctx context.Context, username string) (models.Admin, error) {
	arguments := r.Called(ctx, username)
	return arguments.Get(0).(models.Admin), arguments.Error(1)
}

func (r MysqlAdminRepository) FindAll(ctx context.Context) ([]models.Admin, error) {
	arguments := r.Called(ctx)
	return arguments.Get(0).([]models.Admin), arguments.Error(1)
}

func (r MysqlAdminRepository) FindAdminRoleByID(ctx context.Context, ID int) (string, error) {
	arguments := r.Called(ctx, ID)
	return arguments.String(0), arguments.Error(1)
}
