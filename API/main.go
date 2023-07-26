package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"io"
	"strings"
)

type Word struct {
	Number int    `json:"number"`
	Word   string `json:"word"`
}

func getWordsHandler(w http.ResponseWriter, r *http.Request) {
	var str string

	err := json.NewDecoder(r.Body).Decode(&str)
	fmt.Println("test1")
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Println("test2")
	wor := Word {
		Number: getNumberOfWords(str),
		Word: str,
	}
	fmt.Println("test3")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wor)

	//s = `{"message": "` + getNumberOfWords(str) +  `"}`
	//w.Write([]byte(s)
}

func test(w http.ResponseWriter, r *http.Request) {

	wor := Word {
		Number: getNumberOfWords("str"),
		Word: "str",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wor)
}


func readFile() []string {
	var split []string
	file, err := os.Open("Dracula - Bram Stoker.txt")

	if err != nil {
		fmt.Println("There was an error!")
		return split
	}

	// Closing the file after using it
	defer file.Close()

	// Create a byte slice to hold the file contents
	data := make([]byte, 1024)
	str := ""
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		} 
		if err != nil {
			fmt.Println("File reading error")
			return split
		}
		str += string(data[:n])
	}

	// Get the words to upper case so the code is not case sensitive
	str = strings.ToUpper(str)

	// Split the words into an array so it can be returned
	split = strings.Split(str, " ")

	return split
}

func getNumberOfWords(word string) int {

	// Get the words to upper case so the code is not case sensitive
	word = strings.ToUpper(word)

	count := 0

	// Read the text file and get an array of words
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
	port := ":8080"
	fmt.Println("Server listening on port: ", port)
	router := mux.NewRouter()
		router.HandleFunc("/getNums", test).Methods("GET")
		router.HandleFunc("/getNums", getWordsHandler).Methods("POST")
		
		http.ListenAndServe(port,
			handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			)(router))
}
/*
func main() {
	http.HandleFunc("/getNums", getWordsHandler)
    
	port := ":8080"
	println("Server listening on port: ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
*/