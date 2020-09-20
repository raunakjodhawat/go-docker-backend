// Package app
package Book

import (
	"fmt"
	"github.com/raunakjodhawat/go-docker-backend/internal/DB"
	"net/http"
)

// Book struct
type Book struct {
	BookName string
	Author []string
	ISBN string
	YearOfRelease int
	BookContent []byte
}

func GetAllBooks(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "%v", DB.GetAllBooks())
}

func GetBookById(w http.ResponseWriter, r *http.Request){

}

func RemoveBook(w http.ResponseWriter, r *http.Request){

}
func CreateBook(w http.ResponseWriter, r *http.Request){

}
