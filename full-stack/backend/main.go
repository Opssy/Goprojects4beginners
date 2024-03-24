package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)




type User struct{
	Id    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
func main(){

	//connect to db 

	db, err := sql.Open("postgress", os.Getenv("DATABASE_URL"))

	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	   // create table if it doesn't exist
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
    if err != nil {
        log.Fatal(err)
    }
	//create router
    router := httprouter.New()
	router.GET("/api/go/users", getUsers(db)).Methods("GET")
    router.POST("/api/go/users", createUser(db)).Methods("POST")
    router.GET("/api/go/users/{id}", getUser(db)).Methods("GET")
    router.PUT("/api/go/users/{id}", updateUser(db)).Methods("PUT")
    router.DELETE("/api/go/users/{id}", deleteUser(db)).Methods("DELETE")


	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.Header.Get("Access-Control-Request-Method") != "" {
        // Set CORS headers
        header := w.Header()
        header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
        header.Set("Access-Control-Allow-Origin", "*")
    }

    // Adjust status code to 204
    w.WriteHeader(http.StatusNoContent)
    } )
	// router.GET("/books", controller.GetBooks)
	// router.GET("/books/:id", controller.GetBookById)
	// router.PUT("/books/:id", controller.UpdateBook)
	// router.DELETE("/books/:id", controller.DeleteBook)
	
}