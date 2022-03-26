package repository

import (
	"errors"
	"fmt"
	"homework-2-week-3-HalukErd/core"
	"strings"
)

func GetAllBooks() []core.Book {
	return core.Books
}

func SearchBooksByName(name string) []core.Book {
	var matchedBooks []core.Book
	for _, book := range core.Books {
		if strings.Contains(strings.ToLower(book.Name), strings.ToLower(name)) {
			matchedBooks = append(matchedBooks, book)
		}
	}
	return matchedBooks
}

func SearchBookById(id int64) (core.Book, error) {
	bookResult, _, err := SearchBookAndIndexById(id)
	return bookResult, err
}

func SearchBookAndIndexById(id int64) (core.Book, int, error) {
	for i, book := range core.Books {
		if book.Id == id {
			return book, i, nil
		}
	}
	return core.Book{}, -1, errors.New("Could not found")
}

func DeleteBookById(id int64) error {
	_, i, err := SearchBookAndIndexById(id)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not delete book with id:%d", id))
	}
	core.Books = append(core.Books[:i], core.Books[i+1:]...)
	return nil
}

func UpdateStock(id int64, newQty int) error {
	_, i, err := SearchBookAndIndexById(id)
	if err != nil {
		return err
	}
	core.Books[i].Stock = newQty
	return nil
}
