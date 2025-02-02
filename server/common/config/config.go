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

// func CorsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173/*")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }
