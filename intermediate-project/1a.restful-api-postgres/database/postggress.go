package database

import (
   "fmt"
   "log"
   "gorm.io/driver/postgres"
   "gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Book struct {
   gorm.Model
   Title string
   Author string
}

func DatabaseConnection(){
	host := "localhost"
	port := "5432"
	dbUser := "postgres"
	dbName := "postgres"
	dbPass := "postgres"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, dbUser, dbName, dbPass)


	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	 DB.AutoMigrate(Book{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	fmt.Println("Database connected successfully")
}
