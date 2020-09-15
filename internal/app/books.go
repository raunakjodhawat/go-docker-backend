package app

import "io/ioutil"

// Book struct
type Book struct {
	BookName string
	Author []string
	ISBN string
	YearOfReRelease int
}

// Save stores data into db (POST)
func (b *Book) Save() error {
	filename := b.BookName + ".txt"
	return ioutil.WriteFile(filename, []byte(b.BookName), 0600)
}

// Get all books from db (GET)