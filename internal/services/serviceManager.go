package services

import (
	"errors"
	"fmt"
	"github.com/raunakjodhawat/go-docker-backend/internal/app/book"
	"github.com/raunakjodhawat/go-docker-backend/internal/services/database"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := getAllBooks("mongo")
	fmt.Println(books)
	fmt.Println("hello world")
	fmt.Fprintf(w, "%v", books)
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