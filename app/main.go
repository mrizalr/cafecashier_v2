package main

import (
	"fmt"
	"net/http"

	_adminHandler "github.com/mrizalr/cafecashierpt2/admin/delivery/http"
	_adminRepository "github.com/mrizalr/cafecashierpt2/admin/repository/mysql"
	_adminUsecase "github.com/mrizalr/cafecashierpt2/admin/usecase"
	"github.com/mrizalr/cafecashierpt2/database"
	"github.com/mrizalr/cafecashierpt2/utils"
	"github.com/spf13/viper"
)

func init() {
	utils.InitConfig()
	database.SetDBEnvironment()
}

func main() {
	database.Connect()
	mux := http.NewServeMux()

	adminRepository := _adminRepository.NewMysqlArticleRepository(database.DB())
	adminUcase := _adminUsecase.NewUcaseAdmin(adminRepository)
	_adminHandler.NewAdminHandler(mux, adminUcase)

	http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("server_port")), mux)
}
