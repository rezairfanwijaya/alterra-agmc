package config

import (
	"altera/Day7/pkg/migration"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// load env
	ENV := LoadENV()
	dsn := fmt.Sprintf("root:@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", ENV["host"], ENV["database"])

	// open connection
	var e error
	DB, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if e != nil {
		log.Fatal("Failed connect to database : ", e.Error())
	}

	// migration
	initMigration()
}

func initMigration() {
	err := DB.AutoMigrate(&migration.User{})
	if err != nil {
		log.Fatal("Error migration : ", err.Error())
	}
}

func LoadENV() map[string]string {
	// load file env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error when load env : ", err.Error())
	}

	// get value from env
	host := os.Getenv("HOST")
	database := os.Getenv("DATABASE")
	jwtSecret := os.Getenv("JWT_SECRET")

	return map[string]string{
		"host":      host,
		"database":  database,
		"jwtSecret": jwtSecret,
	}

}
