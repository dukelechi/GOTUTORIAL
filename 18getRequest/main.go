package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Get Request on Go")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // it's your duty to close the body of the request

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	// reading response body
	content, _ := ioutil.ReadAll(response.Body) // remember, content is in the byte format

	// OPTION A; using the string method
	// fmt.Println(string(content)) // that's why we use convert the content using the string method. Also, we need to connect to a local server for this to work. the code is correct

	// alternatively, we can execute this with a 'strings' package
	// remember, string is a data type while strings is a package
	// strings package will give you a whole lot of methods that can be used with a string
	// strings package provides us with a 'Builder'. A builder is used to efficiently build a string using Write methods. It minimizes memory copying. The zero value is ready to use. Do not copy a non-zero Builder.

	// OPTION B; using the string Builder method
	// creating a Builder
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String()) // string method stringifies out whatever data that is saved in the responseString

}
