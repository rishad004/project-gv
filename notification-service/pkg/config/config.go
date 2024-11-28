package config

import (
	"log"

	"github.com/spf13/viper"
)

func Config() error {

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	log.Println("env connected")

	return nil
}
