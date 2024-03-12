package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"main.go/intermediate-project/1a.restful-api-postgres/controller"
	"main.go/intermediate-project/1a.restful-api-postgres/database"
)


func main() {
	fmt.Println("Starting application...")
	database.DatabaseConnection()

    router := httprouter.New()
    router.POST("/books", controller.CreateBook)
	router.GET("/books", controller.GetBooks)
	router.GET("/books/:id", controller.GetBookById)
	router.PUT("/books/:id", controller.UpdateBook)
	router.DELETE("/books/:id", controller.DeleteBook)

	fmt.Println("Listening on port 8080")
}