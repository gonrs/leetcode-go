package runcode

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gonrs/leetcode-go/common/models"
	run "github.com/gonrs/leetcode-go/internal/runCode/lib"
)

// RUN CODE
type RunCodeRequestBody struct {
	ProblemID  uint   `json:"problem_id"`
	Code       string `json:"code"`
	Type       int    `json"type"`
	LanguageID uint   `json:"language_id"`
	// Languages int `json:"languages"`
}

type RunCodeResponse struct {
	Success     bool   `json:"success"`
	Error       string `json:"error"`
	TestOutput  string `json:"test_output"`
	TestInput   string `json:"test_input"`
	TestIndex   int    `json:"test_index"`
	TestsLength int    `json:"tests_length"`
	Output      string `json:"output"`
}

func (h handler) RunCode(ctx *gin.Context) {
	body := RunCodeRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var tests []models.Test

	if err := h.DB.Where("problem_id = ?", body.ProblemID).Where("type = ?", body.Type).Find(&tests).Error; err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if len(tests) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "No test cases found"})
		return
	}
	//
	var language models.LanguageCode

	if result := h.DB.First(&language, body.LanguageID); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	//
	index, out, err := run.Run(language.Language, language.HelpCode, body.Code, tests)

	if err != nil {
		ctx.JSON(http.StatusOK, RunCodeResponse{
			Success:     false,
			Error:       err.Error(),
			TestOutput:  tests[index].OutputForUser,
			TestInput:   tests[index].InputForUser,
			TestIndex:   index,
			TestsLength: len(tests),
			Output:      out,
		})
		return
	}

	ctx.JSON(http.StatusOK, RunCodeResponse{
		Success: true,
		Error:   "",
	})
}

// ADD TESTS
type AddTestsRequestBody struct {
	Input         string `json:"input"`
	Output        string `json"output"`
	InputForUser  string `json:"input_for_user"`
	OutputForUser string `json:"output_for_user"`
	ProblemID     uint   `json:"problem_id"`
	Type          int    `json"type"`
}

func (h handler) AddTests(ctx *gin.Context) {
	body := []AddTestsRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var tests []models.Test
	for _, req := range body {
		tests = append(tests, models.Test{
			Input:         req.Input,
			Output:        req.Output,
			InputForUser:  req.InputForUser,
			OutputForUser: req.OutputForUser,
			ProblemID:     req.ProblemID,
			Type:          req.Type,
		})
	}

	if err := h.DB.Create(&tests).Error; err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Tests added successfully", "tests": tests})
}

// GET TESTS

func (h handler) GetTests(ctx *gin.Context) {
	var tests []models.Test

	if result := h.DB.Find(&tests); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &tests)
}

// DELETE TESTS

func (h handler) DeleteTests(ctx *gin.Context) {
	var ids []int
	if err := ctx.BindJSON(&ids); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if len(ids) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No IDs provided"})
		return
	}

	var tests []models.Test
	if err := h.DB.Where("id IN (?)", ids).Find(&tests).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Some tests not found"})
		return
	}

	if len(tests) != len(ids) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Some tests not found"})
		return
	}

	if err := h.DB.Delete(&tests).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tests"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Test deleted successfully"})
}
