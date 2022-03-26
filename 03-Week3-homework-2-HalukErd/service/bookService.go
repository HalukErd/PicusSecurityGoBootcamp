package service

import (
	"errors"
	"fmt"
	"homework-2-week-3-HalukErd/command"
	"homework-2-week-3-HalukErd/core"
	"homework-2-week-3-HalukErd/repository"
	"strconv"
)

func ListAllBooks(_ command.Params) {
	books := repository.GetAllBooks()
	PrintBooks(books)
}

func SearchBooksAndPrintResult(params command.Params) {
	bookResult, err := getBooksByName(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("---Search Result---")
	PrintBooks(bookResult)
}

func GetBookByIdAndPrintResult(params command.Params) {
	book, err := getBookByID(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(book.ToString())
}

func PrintBooks(books []core.Book) {
	for _, book := range books {
		fmt.Println(book.ToString())
	}
}

func getBooksByName(params command.Params) ([]core.Book, error) {
	param, ok := params["name"]
	if !ok {
		return nil, errors.New("Search by name method needs name param.")
	}
	searchResult := repository.SearchBooksByName(param)
	if len(searchResult) < 1 {
		return nil, errors.New(fmt.Sprintf("Your search: '%s' could not be found\n\n", params["name"]))
	}
	return searchResult, nil
}

func getBookByID(params command.Params) (core.Book, error) {
	idParam, err := getBookIdParam(params)
	if err != nil {
		return core.Book{}, err
	}
	return repository.SearchBookById(idParam)
}

func getBookIdParam(params command.Params) (int64, error) {
	param, ok := params["bookId"]
	if !ok {
		return -1, errors.New("The method needs bookID param.")
	}
	id, err := strconv.Atoi(param)
	if err != nil {
		return -1, err
	}
	return int64(id), nil
}

func getBookQtyParam(params command.Params) (int, error) {
	param, ok := params["bookQty"]
	if !ok {
		return -1, errors.New("The method needs bookID param.")
	}
	qty, err := strconv.Atoi(param)
	if err != nil {
		return -1, err
	}
	return qty, nil
}

func DeleteBookByIdAndPrintResponse(params command.Params) {
	err := deleteBookById(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Book is deleted.")
	fmt.Println("Updated Book List")
	ListAllBooks(nil)
}

func deleteBookById(params command.Params) error {
	idParam, err := getBookIdParam(params)
	if err != nil {
		return errors.New(err.Error())
	}
	return repository.DeleteBookById(idParam)
}

func BuyBookByIdAndQty(params command.Params) {
	err := buyBookAndUpdateStock(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Buy Book process has finished with success.")
	fmt.Println("Updated Book List")
	ListAllBooks(nil)
}

func buyBookAndUpdateStock(params command.Params) error {
	book, err := getBookByID(params)
	if err != nil {
		return err
	}
	qtyParam, err := getBookQtyParam(params)
	if err != nil {
		return err
	}
	if book.Stock == 0 {
		return errors.New("There is no stock.")
	}
	newStock := book.Stock - qtyParam
	if newStock < 0 {
		return errors.New("There is low stock.")
	}
	repository.UpdateStock(book.Id, newStock)
	return nil
}
