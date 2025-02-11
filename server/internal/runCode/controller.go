package runcode

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes := router.Group("/test")
	routes.POST("/run", h.RunCode)
	routes.POST("/add", h.AddTests)
	routes.GET("/get", h.GetTests)
	routes.POST("/delete", h.DeleteTests)
}
