package config

import (
	"fmt"
	"log"
	"os"

	"github.com/JerryJeager/amor-rendezvous-backend/service"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Session *gorm.DB

func GetSession() *gorm.DB{
	return Session
}

func ConnectToDB(){

	// host := os.Getenv("HOST")
	// username := os.Getenv("USER")
	// password := os.Getenv("PASSWORD")
	// port := os.Getenv("PORT")
	// dbName := os.Getenv("DBNAME")

	connectionString := os.Getenv("CONNECTION_STRING")

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbName, port)
    db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil{
		log.Fatal(err)
	}

	db.AutoMigrate(service.User{})

	Session =  db.Session(&gorm.Session{})
	if Session != nil{
		fmt.Println("success: created db session")
	}
}


func LoadEnv(){
    err := godotenv.Load()

	if err != nil{
		fmt.Println(err)
		log.Fatal("failed to load environment variables")
	}
}
