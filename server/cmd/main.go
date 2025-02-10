package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gonrs/leetcode-go/common/db"
	languagecode "github.com/gonrs/leetcode-go/internal/languageCode"
	"github.com/gonrs/leetcode-go/internal/problems"
	runcode "github.com/gonrs/leetcode-go/internal/runCode"
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
	//
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://pt-pt-client.onrender.com"},             // Разрешенные источники
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},                                           // Разрешенные методы
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "Origin", "Content-Type", "Authorization"}, // Разрешенные заголовки
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // Разрешить отправку учетных данных
		MaxAge:           12 * time.Hour, // Время кэширования
	}
	router.Use(cors.New(config))
	//
	problems.RegisterRoutes(router, dbHandler)
	runcode.RegisterRoutes(router, dbHandler)
	languagecode.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	router.Run(port)
}
