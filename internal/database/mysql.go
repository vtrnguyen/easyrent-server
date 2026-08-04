package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"easyrent-server/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect establishes a connection to the MySQL database using GORM and environment variables.
func Connect() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env file not found")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	if err := DB.AutoMigrate(
		&models.Property{},
		&models.PropertyImage{},
		&models.PropertyVideo{},
		&models.PropertyUtility{},
		&models.Utility{},
	); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database successfully")
}
