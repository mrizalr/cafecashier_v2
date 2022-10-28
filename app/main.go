package main

import (
	"github.com/mrizalr/cafecashierpt2/database"
	"github.com/mrizalr/cafecashierpt2/utils"
)

func init() {
	utils.InitConfig()
	database.SetDBEnvironment()
}

func main() {
	database.Connect()
}
