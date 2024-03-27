package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"intermediate-project/2.go-jwt/model"
	"intermediate-project/2.go-jwt/routes"
)

func main() {

	// create a new gin instance

	r := gin.Default()

	    // Load .env file and Create a new connection to the database

		err := godotenv.Load()

		if err != nil{
			 log.Fatal("Error loading .env file")
		}
		config := model.Config{
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName:   os.Getenv("DB_NAME"),
        SSLMode:  os.Getenv("DB_SSLMODE"),
    }

	  // Initialize DB
      model.InitDB(config)

	   // Load the routes
    routes.AuthRoutes(r)

    // Run the server
    r.Run(":8080")
}