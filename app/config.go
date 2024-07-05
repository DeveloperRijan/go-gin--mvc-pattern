package app

import (
	"GoMvcPattern/database"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Boot() {
	//load dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB = database.ConnectDB()

}
