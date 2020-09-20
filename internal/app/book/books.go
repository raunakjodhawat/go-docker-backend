// Package app
package Book

import (
	"io/ioutil"
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

// Save stores data into db (POST, can be changed with PUT)
func (b *Book) SaveBook() error {
	filename := b.BookName + ".txt"
	// 0600 indicates read-write permission
	return ioutil.WriteFile(filename, b.BookContent, 0600)
}

// GetBook, by id
func GetBook(bookName string) (*Book, error) {
	filename := bookName + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Book{BookName: bookName, Author: []string{"a", "b"}, ISBN: "Sample ISBN", YearOfRelease: 2020, BookContent: body}, nil
}

// DeleteBook, deletes a book by id. returns error, if id is not valid or if there's a failure to delete
func (b *Book) DeleteBook() error {
	filename := b.BookName + ".txt"
	return ioutil.WriteFile(filename, []byte(b.BookName), 0600)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request){

}

func GetBookById(w http.ResponseWriter, r *http.Request){

}

func RemoveBook(w http.ResponseWriter, r *http.Request){

}
func CreateBook(w http.ResponseWriter, r *http.Request){

}
