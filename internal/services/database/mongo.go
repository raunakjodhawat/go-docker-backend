package database

import (
	"context"
	"fmt"
	DB "github.com/raunakjodhawat/go-docker-backend/configs"
	"github.com/raunakjodhawat/go-docker-backend/internal/app/book"
	"go.mongodb.org/mongo-driver/bson"
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

func GetAllBooks() []Book.Book {

	client, ctx := getMongoClient()
	defer client.Disconnect(ctx)
	booksDB := client.Database("books")
	booksCollection := booksDB.Collection("books")

	var books []Book.Book
	cursor, err := booksCollection.Find(ctx, bson.M{})
	fmt.Println(cursor)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &books); err != nil {
		panic(err)
	}
	return books
}
