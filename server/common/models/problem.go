package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Title      string `json:"title"`
	Body       string `json:"body"`
	Difficulty int    `json:"difficulty"`
	Code       string `json:"code"`
	HelpCode  string `json:"help_code"`
}
