package main

import (
	"github.com/mrizalr/cafecashierpt2/database"
	"github.com/mrizalr/cafecashierpt2/utils"
)

func Init() {
	utils.InitConfig()
}

func main() {
	database.Connect()
}
