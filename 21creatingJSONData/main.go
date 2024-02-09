package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // alias; Golang renames Name to coursename when the data is printed during JSON encoding
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // won't show the password to whoever is using the API
	Tags     []string `json:"tags,omitempty"` // omitempty; if the value is nil, don't throw that field
}

func main() {
	fmt.Println("Welcome to creating JSON data on Go ")
	EncodeJSON()
}

// ENCODING JSON
// means you want to convert your available data into a valid JSON

func EncodeJSON() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	// package this data as JSON data
	// finalJson, err := json.Marshal(lcoCourses) // hard to read this API with Marshal method. Best to use MarshalIndent method
	// MarshalIndent is like Marshal but applies Indent to format the output. Each JSON element in the output will begin on a new line beginning with prefix followed by one or more copies of indent according to the indentation nesting.
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") // indented based on "\t". No need for prefix ""
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
