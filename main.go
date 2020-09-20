package main

import (
	"github.com/gorilla/mux"
	Book "github.com/raunakjodhawat/go-docker-backend/internal/app/book"
	"log"
	"net/http"
)

const baseURL = "/api/v1"


func booksRouter() *mux.Router {
	const booksURL = "/books/"
	// create a new instance of mux router
	router := mux.NewRouter().StrictSlash(true)

	// GET All Books
	router.HandleFunc(baseURL+booksURL, Book.GetAllBooks).Methods("GET")
	// GET Single book by ID
	router.HandleFunc(baseURL+booksURL+"/{bookId}", Book.GetBookById).Methods("GET")

	// PUT a book
	router.HandleFunc(baseURL+booksURL, Book.CreateBook).Methods("PUT")

	// DELETE Single book by ID
	router.HandleFunc(baseURL+booksURL+"/{bookId}", Book.RemoveBook).Methods("DELETE")

	return router
}

func main() {
	// initialize Books router
	booksRouter := booksRouter()

	// Start serving on 8080 PORT
	log.Fatalln(http.ListenAndServe(":8080", booksRouter))
}
