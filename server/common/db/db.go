package db

import (
	"log"

	"github.com/gonrs/leetcode-go/common/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Problem{})
	db.AutoMigrate(&models.Test{})
	db.AutoMigrate(&models.LanguageCode{})

	return db
}
