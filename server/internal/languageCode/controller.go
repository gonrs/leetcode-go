package languagecode

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

	routes := router.Group("/languagecode")
	routes.POST("/", h.AddLanguage)
	routes.DELETE("/:id", h.DeleteLanguage)
	routes.GET("/:id", h.GetLanguages)
}
