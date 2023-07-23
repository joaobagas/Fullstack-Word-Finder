package main

import (
	"encoding/json"
	"net/http"
	"os"
	"fmt"
	"io"
)

type Word struct {
	Number int    `json:"number"`
	Word   string `json:"word"`
}

var num = Word{Number = 1, Word = "If"}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func readFile() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("There was an error!", err)
		return
	}
	defer file.Close() // it's important to close the file after reading it

	// create a byte slice to hold the file contents
	data := make([]byte, 1024)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}
		fmt.Println("Read", n, "bytes:", string(data[:n]))
	}
}

func getNumberOfWords(word: string) {
	count := 0

	// Get the word
	text := readFile()

	// Foreach word in text file check if it matches the word
	for i := 0 ; i < len(text) ; i++ {
		if (text[i] == word) {
			count++
		}
	}

	return count
	
}

func main() {
	http.HandleFunc("/books", getBooksHandler)

	port := ":8080"
	println("Server listening on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}