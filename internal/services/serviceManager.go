package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/raunakjodhawat/go-docker-backend/internal/app/book"
	"github.com/raunakjodhawat/go-docker-backend/internal/services/database"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := getAllBooks("mongo")
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Error in getting all the books"))
		return
	}
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	books, _ := getAllBooks("mongo")
	fmt.Fprintf(w, "%v", books)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	books, _ := getAllBooks("mongo")
	fmt.Fprintf(w, "%v", books)
}
func RemoveBook(w http.ResponseWriter, r *http.Request) {
	books, _ := getAllBooks("mongo")
	fmt.Fprintf(w, "%v", books)
}

func getAllBooks(dbName string) ([]Book.Book, error) {
	switch dbName {
	case "mongo": {
		return database.GetAllBooks(), nil
	}
	default:
		return nil, errors.New("unsupported DB")
	}

}