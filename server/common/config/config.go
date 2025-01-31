package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `json:"address"`
	DBURL      string `json:"db_url"`
}

func GetConfig() Config {
	viper.AddConfigPath("./common/config/.env")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}
	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err.Error())
	}

	port := viper.Get("PORT")
	dbUrl := viper.Get("DB_URL")

	fmt.Println(port, dbUrl)

	return config
}
