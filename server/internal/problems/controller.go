package problems

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

	routes := router.Group("/problems")
	routes.POST("/", h.AddProblem)
	routes.GET("", h.GetProblems)
	routes.GET("/:id", h.GetProblem)
	routes.PUT("/:id", h.UpdateProblem)
	routes.DELETE("/:id", h.DeleteProblem)
}
