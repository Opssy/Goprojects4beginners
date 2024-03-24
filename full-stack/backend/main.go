package main

import (
	"database/sql"
	"os"
)




type User struct{
	Id    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
func main(){

	//connect to db 

	db, err := sql.Open("postgress", os.Getenv())
}