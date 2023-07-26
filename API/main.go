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

// This struct is used to send the message from the JS frontend to the GO backend
type JSMessage struct {
	Word   string `json:"word"`
}

// This struct is used to send the message from the GO backend to the JS frontend
type GOMessage struct {
	Number int    `json:"number"`
	Word   string `json:"word"`
}

// This function is used to read the file and parse the words into a list of strings
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

	// This goes through every special character and substitutes it by a space
	for _, c := range `".,:;'*[](){}?!¡@"-_/` {
		str = strings.Replace(str, string(c), " ", -1)
	}
	fmt.Println(str)

	// Split the words into an array so it can be returned
	split = strings.Split(str, " ")

	return split
}

// This function is used to count the number of time the word appears in the text file
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

// This function is used to handle the HTTP POST request to get the number of words
func getWordsHandler(w http.ResponseWriter, r *http.Request) {
	var jsMessage JSMessage
	
	err := json.NewDecoder(r.Body).Decode(&jsMessage)
	
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	goMessage := GOMessage {
		Number: getNumberOfWords(jsMessage.Word),
		Word: jsMessage.Word,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(goMessage)
}

func main() {
	port := ":8080"
	fmt.Println("Server listening on port: ", port)
	router := mux.NewRouter()
		router.HandleFunc("/getNums", getWordsHandler).Methods("POST")
		
		http.ListenAndServe(port,
			// This is a fix to the CORS problem
			handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			)(router))
}