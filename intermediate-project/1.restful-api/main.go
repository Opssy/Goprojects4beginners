package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
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

var(
	books = make(map[int]Book)
	booksMutex sync.RWMutex
)

func CreateBook(w http.ResponseWriter, r *http.Request){
	booksMutex.Lock()
	defer booksMutex.Unlock()
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books[book.ID] = book
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusCreated)
}

func GetBook(w http.ResponseWriter, r *http.Request){
	booksMutex.RLock()
	defer booksMutex.RLock()
	if len(books) == 0 {
		log.Println("No books found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	booksMutex.RLock()
	defer booksMutex.RLock()
	if len(books) == 0 {
		log.Println("No books found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	id := r.URL.Query().Get("id")
	book, ok := books[id]
	if !ok {
		log.Println("Book not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
func UpdateBook(w http.ResponseWriter, r *http.Request){
var updateBook Book
	_ = json.NewDecoder(r.Body).Decode(&updateBook)
	booksMutex.Lock()
	defer booksMutex.Unlock()

	_, ok := books[updateBook.ID]
	if !ok {
		log.Println("Book not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	books[updateBook.ID] = updateBook
   
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateBook)
}
func DeleteBook(w http.ResponseWriter, r *http.Request){
	booksMutex.Lock()
	defer booksMutex.Unlock()

	id := r.URL.Query().Get("id")
	_, ok := books[id]
	if !ok {
		log.Println("Book not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	intID, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Invalid ID format")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	delete(books, intID)
	w.WriteHeader(http.StatusOK)
		
}
func main(){
	books[1] = Book{ID: 1, Title: "Harry Potter", Author: "JK Rowling", Publisher: &Company{Name: "Bloomsbury", Address: "London", Phone: "1234567890"}}
	books[2] = Book{ID: 2, Title: "Lord of the Rings", Author: "Tolkien", Publisher: &Company{Name: "Allen & Unwin", Address: "New York", Phone: "0987654321"}}

	http.HandleFunc("/books", CreateBook)
	http.HandleFunc("/books", GetBook)
	http.HandleFunc("/books/id", GetBookById)
	http.HandleFunc("/books/id", UpdateBook)
	http.HandleFunc("/books/id", DeleteBook)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
