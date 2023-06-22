package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * DB is a global variable that holds the database connection.
 * It is initialized in the main function.
 */

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(localhost:3306)/go_restapi_gin?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Product{})

	DB = database
}
