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
	fmt.Println("Welcome to JSON ")
	// EncodeJSON()
	DecodeJSON()
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

// CONSUMING JSON Data; same thing as decoding JSON Data
func DecodeJSON() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in", 
		"tags": ["web-dev","js"]
	}
	`)

	// OPTION A; with struct
	// making lcoCourses as a type of course (struct)
	var lcoCourses course

	// checking the validity of the data we copied from the web
	checkValid := json.Valid(jsonDataFromWeb) // Valid reports whether data is a valid JSON encoding. Gives a bool value
	if checkValid {
		fmt.Println("JSON IS VALID")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses) // in order to print these interfaces, we need to have a special synthax "%#v\n"
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	// OPTION B; with map
	// some cases were you just want to add data to key-value without creating a struct
	var myOnlineData map[string]interface{} // the map will be a type of string, and since we are not sure it's only string will be the type of data that will be pulled from the web, we use an interface
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	// OPTION C; with a for loop
	// looping through the key value pairs
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	} // note that order is not guaranteed on the key-value pairs
}
