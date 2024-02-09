package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// formdata
	data := url.Values{}             // values{} creates an empty key-value pairs
	data.Add("firstname", "Kelechi") // adds data to Values{}
	data.Add("lastname", "Duru")
	data.Add("email", "hitesh@go.dev")

	// making Post Form Request
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	// closing the body after reading

	// reading the content
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content)) // needs to be connected to a server to work // prints out the added Values
}

func main() {
	fmt.Println("Welcome to Sending Form data in Go")
	PerformPostFormRequest()
}
