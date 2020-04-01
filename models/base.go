package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) //Build connection string

	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	// defer conn.Close()

	if err != nil {
		fmt.Printf("Error while connecting to db : %s ", err)
		os.Exit(1)
	}

	db = conn
	db.Debug().AutoMigrate(&User{}) //Database migration

}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
