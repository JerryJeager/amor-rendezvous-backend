package config

import (
	"fmt"
	"log"

	// "os"

	"github.com/JerryJeager/amor-rendezvous-backend/service"
	// "github.com/joho/godotenv"
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
	// port := os.Getenv("DBPORT")
	// dbName := os.Getenv("DBNAME")

	// connectionString := os.Getenv("CONNECTION_STRING")

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbName, port)

	// dsn := "host=localhost user=postgres password=chidiebere823A dbname=amor_rendezvous port=5432 sslmode=disable"

    db, err := gorm.Open(postgres.Open("postgres://postgres.qnwmvvwjqgonoromlfwm:chidiebere823A@aws-0-eu-central-1.pooler.supabase.com:5432/postgres"), &gorm.Config{})
	
    // db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal(err)
	}

	db.AutoMigrate(service.User{})
	db.AutoMigrate(service.Wedding{})

	Session =  db.Session(&gorm.Session{})
	if Session != nil{
		fmt.Println("success: created db session")
	}
}


// func LoadEnv(){
//     err := godotenv.Load()

// 	if err != nil{
// 		fmt.Println(err)
// 		log.Fatal("failed to load environment variables")
// 	}
// }
