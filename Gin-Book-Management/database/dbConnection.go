package database

import (
	"fmt"
	"gin-book-management/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var err error
var urlDsn = "root:@tcp(127.0.0.1:3306)/gin_book_management?charset=utf8mb4&parseTime=True&loc=Local"

func DatabaseMigration() {
	Database, err = gorm.Open(mysql.Open(urlDsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("Cannot connect to DB")
	}

	log.Println("Connected to Database...")
	Database.AutoMigrate(models.User{})
	Database.AutoMigrate(models.Categories{})
	Database.AutoMigrate(models.Books{})
	Database.AutoMigrate(models.Author{})

}
