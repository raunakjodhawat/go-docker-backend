package DB

import (
	"errors"
	"github.com/raunakjodhawat/go-docker-backend/internal/DB/Mongo"
	"github.com/raunakjodhawat/go-docker-backend/internal/app/book"
	)

func GetAllBooks(dbName string) ([]Book.Book, error) {
	switch dbName {
	case "Mongo": {
		return Mongo.MongoGetAllBooks(), nil
	}
	default:
		return nil, errors.New("unsupported DB")
	}

}