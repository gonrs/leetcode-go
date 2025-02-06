package runcode

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gonrs/leetcode-go/common/models"
	run "github.com/gonrs/leetcode-go/internal/runCode/lib"
)

// RUN CODE
type RunCodeRequestBody struct {
	ProblemID uint   `json:"problemId"`
	Code      string `json:"code"`
	Type      int    `json"type"`
	// Languages int `json:"languages"`
}

type RunCodeResponse struct {
	Success  bool   `json:"success"`
	LastTest int    `json:"last_test"`
	Error    string `json:"error"`
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
	var problem models.Problem

	if result := h.DB.First(&problem, body.ProblemID); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	//
	res, err := run.Run(problem.HelpCode, body.Code, tests)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, RunCodeResponse{
			Success:  false,
			LastTest: res,
			Error:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, RunCodeResponse{
		Success:  true,
		LastTest: -1,
		Error:    "",
	})
}

// ADD TESTS
type AddTestsRequestBody struct {
	Input     string `json:"input"`
	Output    string `json"output"`
	ProblemID uint   `json:"problem_id"`
	Type      int    `json"type"`
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
			Input:     req.Input,
			Output:    req.Output,
			ProblemID: req.ProblemID,
			Type:      req.Type,
		})
	}

	if err := h.DB.Create(&tests).Error; err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Tests added successfully", "tests": tests})
}
