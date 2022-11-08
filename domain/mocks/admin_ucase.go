package mocks

import (
	"context"

	"github.com/mrizalr/cafecashierpt2/domain"
	"github.com/mrizalr/cafecashierpt2/models"
	"github.com/stretchr/testify/mock"
)

type AdminUcase struct {
	mock.Mock
}

func (u *AdminUcase) Add(ctx context.Context, req *models.CreateNewAdminRequest) (res domain.Admin, err error) {
	arguments := u.Called(ctx, req)
	return arguments.Get(0).(domain.Admin), arguments.Error(1)
}
