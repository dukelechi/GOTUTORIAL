package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Post Request Function
func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake body payload
	requestBody := strings.NewReader(`
		{
			"coursename" : "Let's go with golang"
			"price" : 0,
			"platform" : "learncodeonline.in"
		}
	`)

	// sending request
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	// close the body after reading
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	// print out the data
	fmt.Println(string(content)) // needs to connect to server to work. Prints out {"coursename" : "Let's go with golang", "price" : 0,"platform" : "learncodeonline.in"}

}

func main() {
	fmt.Println("Welcome to Post Request on Go")
	PerformPostJsonRequest()
}
