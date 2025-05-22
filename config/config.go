package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func Initiator() {
	log.Println("Initiating Configuration....")

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	log.Println("Successfully read config file")
}
