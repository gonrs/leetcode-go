package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Title      string `json:"title"`
	Body       string `json:"body"`
	Solution   string `json:"solution"`
	Difficulty int    `json:"difficulty"`
}
