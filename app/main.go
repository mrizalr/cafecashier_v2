package main

import (
	"context"

	_adminRepository "github.com/mrizalr/cafecashierpt2/admin/repository/mysql"
	_adminUsecase "github.com/mrizalr/cafecashierpt2/admin/usecase"
	"github.com/mrizalr/cafecashierpt2/database"
	"github.com/mrizalr/cafecashierpt2/models"
	"github.com/mrizalr/cafecashierpt2/utils"
)

func init() {
	utils.InitConfig()
	database.SetDBEnvironment()
}

func main() {
	database.Connect()

	adminRepository := _adminRepository.NewMysqlArticleRepository(database.DB())
	adminUcase := _adminUsecase.NewUcaseAdmin(adminRepository)

	request := models.CreateNewAdminRequest{
		Username: "owner",
		Password: "owner123",
	}

	adminUcase.Add(context.Background(), request)
}
