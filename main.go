package main

import (
	Book "./internal/app/book"
	"fmt"
)

func main() {
	b := Book.Book{Author: []string{"Raunak Jodhawat"}, BookName: "RJ", ISBN: "HELLO", YearOfRelease: 2020, BookContent: []byte("Hello, this is my first book") }
	b.SaveBook()

	b1, err := Book.GetBook("RJ")
	if err == nil {
		fmt.Println(string(b1.BookContent))
	}
}
