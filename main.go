package main

import (
	Book "./internal/app/book"
	"fmt"
	"log"
	"net/http"
)

const baseUrl = "/api/v1"
const booksUrl = "/books/"
func booksHandler(w http.ResponseWriter, r *http.Request) {
	bookName := r.URL.Path[len(baseUrl+booksUrl):]
	book, err := Book.GetBook(bookName)
	if err == nil {
		fmt.Fprintf(w, "Found the Book, Author: %s, %s, %s", book.Author, book.ISBN, book.BookContent)
	} else {
		fmt.Fprintf(w, "Not able to find %s, book", r.URL.Path[len(baseUrl+booksUrl):])
	}

}
func main() {

	http.HandleFunc(baseUrl+booksUrl, booksHandler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

//