package models

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	Input     string `json:"input"`
	Output    string `json:"output"`
	ProblemID uint   `json:"problem_id"`
	Type      int    `json:"type"`
}
