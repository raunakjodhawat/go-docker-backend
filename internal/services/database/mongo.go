package database

import (
	"context"
	DB "github.com/raunakjodhawat/go-docker-backend/configs"
	"github.com/raunakjodhawat/go-docker-backend/internal/app/book"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)


func getMongoClient() (*mongo.Client, context.Context) {
	mongoConfig := DB.MongoConfig()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConfig.GetConnectionString()))
	if err != nil {
		panic(err)
	}

	return client, ctx
}

func GetAllBooks() ([]Book.Book, error) {

	client, ctx := getMongoClient()
	defer client.Disconnect(ctx)
	booksDB := client.Database("books")
	booksCollection := booksDB.Collection("books")

	var books []Book.Book
	cursor, err := booksCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &books); err != nil {
		return nil, err
	}
	return books, err
}

func GetBookById(_id primitive.ObjectID) (Book.Book, error) {
	client, ctx := getMongoClient()
	defer client.Disconnect(ctx)
	booksDB := client.Database("books")
	booksCollection := booksDB.Collection("books")

	var book Book.Book
	err := booksCollection.FindOne(ctx, bson.D{{"_id", _id}}).Decode(&book)
	if err != nil{
		return Book.Book{}, err
	}
	return book, nil
}