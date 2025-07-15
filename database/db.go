package database

import (
	"fmt"
	"log"
	"task-manager/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// MySQL connection parameters
	username := "root" // Default MySQL username
	password := "root" // Change this to your MySQL password
	host := "localhost"
	port := "3306"
	dbname := "task_manager"

	// MySQL connection URL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbname,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&models.Task{}, &models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Database connected successfully")
}
