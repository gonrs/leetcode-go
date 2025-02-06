package problems

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gonrs/leetcode-go/common/models"
)

// FUNCTION: ADD PROBLEM

type AddProblemRequestBody struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	Difficulty int    `json:"difficulty"`
	Code       string `json:"code"`
	HelpCode   string `json:"help_code"`
}

func (h handler) AddProblem(ctx *gin.Context) {
	body := AddProblemRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var problem models.Problem
	//
	problem.Title = body.Title
	problem.Body = body.Body
	problem.Difficulty = body.Difficulty
	problem.Code = body.Code
	problem.HelpCode = body.HelpCode
	//
	if result := h.DB.Create(&problem); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &problem)
}

// FUNCTION: GET ALL PROBLEMS
func (h handler) GetProblems(ctx *gin.Context) {
	var problems []models.Problem
	start, end := ctx.DefaultQuery("from", "0"), ctx.DefaultQuery("to", "3")
	s, err := strconv.Atoi(start)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	e, err := strconv.Atoi(end)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	if result := h.DB.Offset(s).Limit(e - s).Find(&problems); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &problems)
}

// FUNCTION: GET PROBLEM
type GetProblemResponse struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	Difficulty int    `json:"difficulty"`
	Code       string `json:"code"`
}

func (h handler) GetProblem(ctx *gin.Context) {
	id := ctx.Param("id")

	var problem models.Problem

	if result := h.DB.First(&problem, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, GetProblemResponse{
		Title:      problem.Title,
		Body:       problem.Body,
		Difficulty: problem.Difficulty,
		Code:       problem.Code,
	})
}

// // FUNCTION: GET PROBLEM

// func (h handler) GetProblemHelpCode(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	var problem models.Problem

// 	if result := h.DB.First(&problem, id); result.Error != nil {
// 		ctx.AbortWithError(http.StatusNotFound, result.Error)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, problem.HelpCode)
// }

// FUNCTION: UPDATE PROBLEM
func (h handler) UpdateProblem(ctx *gin.Context) {
	id := ctx.Param("id")
	body := AddProblemRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var problem models.Problem

	if result := h.DB.First(&problem, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	//
	problem.Title = body.Title
	problem.Body = body.Body
	problem.Difficulty = body.Difficulty
	problem.Code = body.Code
	problem.HelpCode = body.HelpCode
	//

	h.DB.Save(&problem)
	ctx.JSON(http.StatusOK, &problem)
}

// FUNCTION: DELETE PROBLEM
func (h handler) DeleteProblem(ctx *gin.Context) {
	id := ctx.Param("id")

	var problem models.Problem

	if result := h.DB.First(&problem, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&problem)

	ctx.Status(http.StatusOK)
}
