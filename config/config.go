package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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
    p, err := os.Executable()
    if err != nil {
        log.Fatal(err)
    }

    // Extract the directory containing the executable
    p = filepath.Dir(p)

    // Load the .env file
    err = godotenv.Load(filepath.Join(p, ".env"))
    if err != nil {
        panic("Error loading .env file")
    }

}
