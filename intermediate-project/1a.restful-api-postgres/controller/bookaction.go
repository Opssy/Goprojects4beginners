package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"main.go/intermediate-project/1a.restful-api-postgres/database"
)

func CreateBook(w http.ResponseWriter, r *http.Request , _ httprouter.Params) {
	var book *database.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	res := database.DB.Create(book)
	if res != nil {
		log.Println(res)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Book created successfully"))

}

func GetBooks(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
	var books []database.Book
	database.DB.Find(&books)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get Books"))
}
func GetBookById(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
	id := r.URL.Query().Get("id")
	var book *database.Book
	res := database.DB.First(&book, id)
	if res.RowsAffected == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get Book By Id"))
}
func UpdateBook(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
	id := r.URL.Query().Get("id")
	var book *database.Book
	res := database.DB.First(&book, id)
	if res.RowsAffected == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	res = database.DB.Save(book)
	if res.RowsAffected == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book updated successfully"))
}

func DeleteBook(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
	id := r.URL.Query().Get("id")
	var book *database.Book
	res := database.DB.First(&book, id)
	if res.RowsAffected == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	res = database.DB.Delete(book)
	if res.RowsAffected == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book deleted successfully"))
}