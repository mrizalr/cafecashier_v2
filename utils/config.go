package utils

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigType("json")
	wd, _ := os.Getwd()
	viper.AddConfigPath(wd)
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error when read config file : " + err.Error())
	}

	log.Println("Success init config.json")
}
