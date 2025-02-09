package models

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	InputForUser string `json:"input_for_user"`
	OutputForUser string `json:"output_for_user"`
	Input        string `json:"input"`
	Output       string `json:"output"`
	ProblemID    uint   `json:"problem_id"`
	Type         int    `json:"type"`
}
