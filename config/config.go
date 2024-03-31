package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Session *gorm.DB

func GetSession() *gorm.DB{
	return Session
}

func ConnectToDB(){

	host := os.Getenv("HOST")
	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	port := os.Getenv("PORT")
	dbName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbName, port)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal(err)
	}

	Session = db
	if Session != nil{
		fmt.Println("success: created db session")
	}
}


func LoadEnv(){
    err := godotenv.Load()

	if err != nil{
		log.Fatal("failed to load environment variables")
	}
}