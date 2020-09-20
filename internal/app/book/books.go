// Package app
package Book

import (
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

}

func GetBookById(w http.ResponseWriter, r *http.Request){

}

func RemoveBook(w http.ResponseWriter, r *http.Request){

}
func CreateBook(w http.ResponseWriter, r *http.Request){

}
