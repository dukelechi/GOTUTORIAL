package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Module in Go")
	greeter()
	r := mux.NewRouter() // copied from gorilla package
	// r.HandleFunc("/", HomeHandler) // copied from gorilla pkg; r is on a func HandleFunc and a route "/", and it defines the HomeHandler method. It's edited to the code below
	r.HandleFunc("/", serveHome).Methods("GET") // dot Methods("GET") means you are only serving the get method

	// RUNNING IT AS A SEVER
	log.Fatal(http.ListenAndServe(":5000", r)) // log pkg has the Fatal method which auto logs that value if something fails
}

// function greater
func greeter() {
	fmt.Println("Hey there, mod users!")
}

// w for response writter, r for request
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series on YouTube<h1>"))
}
