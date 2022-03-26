package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Book struct {
	name string
}

type cmd struct {
	key  string
	task func(params params)
}

type params map[string]string

type cmdReq struct {
	cmdKey string
	params params
}

var commands = make(map[string]cmd)

func init() {
	populateCommands()
}

func populateCommands() {
	commands["list"] = cmd{"list", printAllBooks}
	commands["search"] = cmd{"search", printResultOfSearchBook}
}

func main() {
	cmdReq := handleReadArgs()
	cmd := commands[cmdReq.cmdKey]
	cmd.task(cmdReq.params)
}

func handleReadArgs() cmdReq {
	cmd, err := getCmdAndArgs()
	if err != nil {
		fmt.Println("You need to enter at least a cmdReq.")
	}
	return cmd
}

func getCmdAndArgs() (cmdReq, error) {
	args := os.Args
	var cmd cmdReq
	if len(args) < 2 {
		return cmd, errors.New("not enough args")
	}
	cmd.cmdKey = args[1]
	if len(args) > 2 {
		paramWords := args[2:]
		param := make(map[string]string)
		param["name"] = strings.Join(paramWords, " ")
		cmd.params = param
	}

	return cmd, nil
}

func printAllBooks(_ params) {
	books := listAllBooks()

	for _, book := range books {
		fmt.Println(book.name)
	}
}

func printResultOfSearchBook(params params) {
	bookResult := searchByName(params)
	if bookResult == nil {
		fmt.Printf("Your search: '%s' could not be found\n", params["name"])
		return
	}
	fmt.Printf("You found %s\n", bookResult[0].name)
}

func listAllBooks() []Book {
	file := readFile("books")
	defer file.Close()

	var books []Book

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		books = append(books, Book{name: line})
	}
	return books
}

func searchByName(params params) []Book {
	param, ok := params["name"]
	if !ok {
		panic("Search by name method needs name param.")
	}

	file := readFile("books")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matched := strings.ToLower(line) == strings.ToLower(param)
		if matched {

			return []Book{{name: line}}
		}
	}

	return nil
}

func readFile(fName string) *os.File {
	file, err := os.Open(fName)
	if err != nil {
		log.Panicf("%s could not be opened!", fName)
	}
	return file
}
