package services

import (
	"fmt"
	"os"

	"github.com/C-m3-Codin/buffalo_gin_bench/gin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// NewDatabase : intializes and returns mysql db
func NewDatabase() (*gorm.DB, error) {

	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
	ENV := os.Getenv("ENVIRONMENT")

	if ENV == "" {
		fmt.Println("No Env")
		USER = "cp"
		PASS = "cp_pass"
		HOST = "localhost"
		// PORT = 3306
		DBNAME = "testDb"
	}

	// fmt.Println(env)
	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS,
		HOST, DBNAME)
	fmt.Println(URL)
	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Set the logger to print info-level logs.
	})

	mysqldb, err := db.DB()
	mysqldb.SetMaxIdleConns(10)
	mysqldb.SetMaxOpenConns(100)

	if err != nil {
		fmt.Println("\n\n\n ", err)
		panic("Failed to connect to database!")
		return nil, err

	}
	// Create the necessary tables in the database.
	db.AutoMigrate(models.User{})
	DB = db
	if err != nil {
		// panic("Failed to connect to database!")
		return nil, err

	}
	fmt.Println("Database connection established")
	return db, nil

}

// USER := os.Getenv("DB_USER")
// PASS := os.Getenv("DB_PASSWORD")
// HOST := os.Getenv("DB_HOST")
// DBNAME := os.Getenv("DB_NAME")
