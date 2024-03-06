package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct{
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"` 
	Publisher  *Company `json:"publisher"`
}
type Company struct{
	Name string `json:"name"`
	Address string `json:"address"`
	Phone string `json:"phone"`
}

var books []Book

func CreateBook(w http.ResponseWriter, r *http.Request){
	var book Book 
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusCreated)
}

func GetBook(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(books)
	for _, book := range books{
	    json.NewEncoder(w).Encode(book)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
func GetBookById(w http.ResponseWriter, r *http.Request){
	for _, book := range books{
		if book.ID == 1{
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	w.WriteHeader(http.StatusOK)

}
func UpdateBook(w http.ResponseWriter, r *http.Request){
	var book Book 
	_ = json.NewDecoder(r.Body).Decode(&book)
	for i, b := range books{
		if b.ID == book.ID{
			books[i] = book
			json.NewEncoder(w).Encode(book)
			return
		}
		w.WriteHeader(http.StatusNotFound)

	}
}
func DeleteBook(w http.ResponseWriter, r *http.Request){
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	for i, b := range books{
		if b.ID == book.ID{
			books = append(books[:i], books[i+1:]...)
			json.NewEncoder(w).Encode(book)
			return
		}
		w.WriteHeader(http.StatusNotFound)

	}
	w.WriteHeader(http.StatusNotFound)
}
func main(){
	books = append(books, Book{ID: 1, Title: "Harry Potter", Author: "JK Rowling", Publisher: &Company{Name: "Bloomsbury", Address: "London", Phone: "1234567890"}})
	books = append(books, Book{ID: 2, Title: "Lord of the Rings", Author: "Tolkien", Publisher: &Company{Name: "Allen & Unwin", Address: "New York", Phone: "0987654321"}})
	http.HandleFunc("/books", CreateBook)
	http.HandleFunc("/books", GetBook)
	http.HandleFunc("/books/1", GetBookById)
	http.HandleFunc("/books/1", UpdateBook)
	http.HandleFunc("/books/1", DeleteBook)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
