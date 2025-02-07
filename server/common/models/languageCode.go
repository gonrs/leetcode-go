package models

import "gorm.io/gorm"

type LanguageCode struct {
	gorm.Model
	ProblemID uint   `json:"problem_id"`
	Language  string `json:"language"`
	StartCode string `json:"start_code"`
	HelpCode  string `json:"help_code"`
}
