package utils

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error when read config file : " + err.Error())
	}

	log.Println("Success init config.json")
}
