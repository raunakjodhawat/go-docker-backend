package main

import (
	"github.com/gorilla/mux"
	"github.com/raunakjodhawat/go-docker-backend/internal/services"
	"log"
	"net/http"
)

const baseURL = "/api/v1"


func booksRouter() *mux.Router {
	const booksURL = "/books/"
	// create a new instance of mux router
	router := mux.NewRouter().StrictSlash(true)

	// GET All Books
	router.HandleFunc(baseURL+booksURL, services.GetAllBooks).Methods("GET")
	// GET Single book by ID
	router.HandleFunc(baseURL+booksURL+"/{bookId}", services.GetBookById).Methods("GET")

	// PUT a book
	router.HandleFunc(baseURL+booksURL, services.CreateBook).Methods("PUT")

	// DELETE Single book by ID
	router.HandleFunc(baseURL+booksURL+"/{bookId}", services.RemoveBook).Methods("DELETE")

	return router
}

func main() {
	// initialize Books router
	booksRouter := booksRouter()

	// Start serving on 8080 PORT
	log.Fatalln(http.ListenAndServe(":8080", booksRouter))
}
