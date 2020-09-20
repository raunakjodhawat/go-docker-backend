package main

import (
	Book "./internal/app/book"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const baseUrl = "/api/v1"
const booksUrl = "/books/"

func booksRouter() *mux.Router {
	// create a new instance of mux router
	router := mux.NewRouter().StrictSlash(true)

	// GET All Books
	router.HandleFunc(baseUrl+booksUrl, Book.GetAllBooks).Methods("GET")
	// GET Single book by ID
	router.HandleFunc(baseUrl+booksUrl+"/{bookId}", Book.GetBookById).Methods("GET")

	// PUT a book
	router.HandleFunc(baseUrl+booksUrl, Book.CreateBook).Methods("PUT")

	// DELETE Single book by ID
	router.HandleFunc(baseUrl+booksUrl+"/{bookId}", Book.RemoveBook).Methods("DELETE")

	return router
}


func main() {
	// initialize Books router
	booksRouter := booksRouter()
	// Start serving on 8080 PORT
	log.Fatalln(http.ListenAndServe(":8080", booksRouter))
}