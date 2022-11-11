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

	_ "github.com/mrizalr/cafecashierpt2/docs"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

// @title Swagger Cafe Cashier API
// @version 1.0
// @description this is a portofolio project of mrizalr

// @contact.name API SUPPORT
// @contact.url github.com/mrizalr
// @contact email muhammadrizal2252@gmail.com

// @host localhost:8080

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

	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("server_port")), mux)
}
