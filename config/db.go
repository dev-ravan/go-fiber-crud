package config

import (
	"fmt"
	"os"

	"example.com/go-crud/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnection() {
	godotenv.Load()
	dbhost := os.Getenv("MYSQL_HOST")
	dbuser := os.Getenv("MYSQL_USER")
	dbpassword := os.Getenv("MYSQL_PWD")
	dbname := os.Getenv("MYSQL_DBNAME")

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbname)
	DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})

	// Connect db
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database :(")
	}
	println("DB Connected successfully..!")
	DB.AutoMigrate(&models.User{})
}
