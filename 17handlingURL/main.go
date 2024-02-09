package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000/learn?coursename=reactjs&paypmentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to Handling URLs in Golang")
	fmt.Println(myURL)

	// parsing/extracting content (in the context of URL)
	result, _ := url.Parse(myURL) // from URL library

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port()) // port here is a method, not a property
	fmt.Println(result.RawQuery)

	qParams := result.Query()                                     // stores data in key-value pairs
	fmt.Printf("The type of query parameters are: %T\n", qParams) // url.Values = key value pairs

	fmt.Println(qParams["coursename"]) // reactjs
	fmt.Println(qParams["paymentid"])  // ghbj456ghb

	for _, val := range qParams {
		fmt.Println("Parameter is: ", val)
	}

	// CONSTRUCTING A NEW URL
	// always pass on the exact data with a reference "&url"
	// the parameters/properties/keys will have to be exactly as it's written below
	partsOfurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherURL := partsOfurl.String() // using String() prints partsOfUrl as a strings
	fmt.Println(anotherURL)

}
