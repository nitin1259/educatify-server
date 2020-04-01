package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB //database

func init() {

	err := godotenv.Load() //Load .env file

	if err != nil {
		fmt.Println(err)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) //Build connection string

	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)

	if err != nil {
		fmt.Print(err)
	}

	db = conn
	// db.Debug().AutoMigrate(&Account{}, &Contact{}) //Database migration

}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
