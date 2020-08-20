package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func Load() {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "" {
		appEnv = "development"
	}

	viper.SetConfigName(appEnv)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}
