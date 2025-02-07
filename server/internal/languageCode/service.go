package languagecode

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gonrs/leetcode-go/common/models"
)

// ADD LANGUAGE
type AddLanguageRequest struct {
	ProblemID uint   `json:"problem_id"`
	Language  string `json:"language"`
	StartCode string `json:"start_code"`
	HelpCode  string `json:"help_code"`
}

func (h handler) AddLanguage(ctx *gin.Context) {
	var body AddLanguageRequest

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var language models.LanguageCode
	language.ProblemID = body.ProblemID
	language.Language = body.Language
	language.StartCode = body.StartCode
	language.HelpCode = body.HelpCode

	if result := h.DB.Create(&language); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &language)
}

// DELETE LANGUAGE
func (h handler) DeleteLanguage(ctx *gin.Context) {
	id := ctx.Param("id")

	var language models.LanguageCode

	if result := h.DB.First(&language, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&language)

	ctx.Status(http.StatusOK)
}

// FUNCTION: GET LANGUAGE

func (h handler) GetLanguages(ctx *gin.Context) {
	id := ctx.Param("id")

	var language []models.LanguageCode

	if result := h.DB.Where("problem_id = ?", id).Find(&language); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, language)
}
