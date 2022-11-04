package mocks

import (
	"context"

	"github.com/mrizalr/cafecashierpt2/domain"
	"github.com/stretchr/testify/mock"
)

type MysqlAdminRepository struct {
	mock.Mock
}

func (r MysqlAdminRepository) Add(ctx context.Context, admin *domain.Admin) (int64, error) {
	arguments := r.Called(ctx, admin)
	return int64(arguments.Int(0)), arguments.Error(1)
}

func (r MysqlAdminRepository) FindByID(ctx context.Context, id int) (domain.Admin, error) {
	arguments := r.Called(ctx, id)
	return arguments.Get(0).(domain.Admin), arguments.Error(1)
}
