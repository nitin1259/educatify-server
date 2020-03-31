package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file

	if e != nil {
		fmt.Print(e)
	}

	username := e.Getenv("db_user")
	password := e.Getenv("db_pass")
	dbName := e.Getenv("db_name")
	dbHost := e.Getenv("db_host")
	dbPort := e.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbhodbHost, dbPort, username, dbName, password) //Build connection string

	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)

	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{}) //Database migration

}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
