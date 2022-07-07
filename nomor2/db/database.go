package db

import (
	"os"
	"test-efishery/model"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDatabase() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func Migration() {
	DB.AutoMigrate(&model.FormModel{})
}
