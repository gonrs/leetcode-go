package main

import (
	"github.com/gonrs/leetcode-go/common/db"
	"github.com/gonrs/leetcode-go/internal/problems"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./common/config/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	//

	router := gin.Default()
	dbHandler := db.Init(dbUrl)

	problems.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	router.Run(port)
}
