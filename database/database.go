package database

import (
	"fmt"
	"log"
	"omh-simple-app/models"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Singleton DB
var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("MYSQL_USERNAME") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_HOSTNAME") + ":" + os.Getenv("MYSQL_PORT") + ")/" + os.Getenv("MYSQL_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to DB")
	}

	DB = db

	// Drop tables
	DB.Migrator().DropTable(&models.Country{}, &models.Property{})

	// Perform DB migrations
	fmt.Println("Performing migrations")
	DB.AutoMigrate(&models.Country{}, &models.Property{})

	// Seed initial data
	fmt.Println("Seeding data")
	seedData()
}

func seedData() {
	countriesStr := "Singapore, Malaysia, Philippines"
	countries := strings.Split(countriesStr, ",")

	for _, val := range countries {
		c := models.Country{
			Name: strings.TrimSpace(val),
		}

		DB.Create(&c)
	}
}
