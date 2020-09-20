package DB

import (
	"github.com/raunakjodhawat/go-docker-backend/internal/DB/Mongo"
	"github.com/raunakjodhawat/go-docker-backend/internal/app/book"
	)
func GetAllBooks() [] Book.Book {
	return Mongo.MongoGetAllBooks()
}