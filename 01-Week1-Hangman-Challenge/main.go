package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var inputReader = bufio.NewReader(os.Stdin)

var dictionary = []string{
	"masa",
	"sandalye",
	"elma",
	"almanya",
}

const letterRegexp = "[a-z]"

var word string
var emptyWord []byte
var m = make(map[rune][]int)

func main() {

	// Preparation
	i := 0
	clear()
	defineNewWord()

	for {

		// Show status
		clear()
		showGallows(i)
		fmt.Println(string(emptyWord))

		// Get input
		fmt.Printf("Enter a rune: ")
		line, _, _ := inputReader.ReadLine()
		ln := string(line)
		matched, _ := regexp.MatchString(letterRegexp, ln)
		if len(ln) != 1 || !matched {
			fmt.Println("You must enter a letter.")
			time.Sleep(time.Second)
			continue
		}
		
		ch := ln[0]
		indexes, ok := m[rune(ch)]
		if ok {
			for _, index := range indexes {
				emptyWord[index] = ch
			}
			delete(m, rune(ch))
			if len(m) == 0 {
				wellDone(&i)
				defineNewWord()
			}
			continue
		}

		i++
		if i >= 9 {
			fmt.Println("The answer was ...")
			time.Sleep(time.Second * 2)
			fmt.Println(word)
			defineNewWord()
			gameOver(&i)
		}
	}
}

func defineNewWord() {
	word = newRandomWord()
	emptyWord = []byte(strings.Repeat("_", len(word)))
	for i, ch := range word {
		m[ch] = append(m[ch], i)
	}
}

func wellDone(i *int) {
	clear()
	fmt.Println("Well Done")
	*i = 0
	time.Sleep(time.Second * 3)
}

func gameOver(i *int) {
	clear()
	showGallows(9)
	fmt.Println("GAME OVER!")
	time.Sleep(time.Second * 3)
	*i = 0
}

func showGallows(i int) {
	hangPos := "hangman" + strconv.Itoa(i)
	file, _ := os.Open("states/" + hangPos)
	io.Copy(os.Stdout, file)
}

func newRandomWord() string {
	rand.Seed(time.Now().UnixNano())
	return dictionary[rand.Intn(len(dictionary))]
}

func clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
