package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Hi, thanks for rating,", input)

	// let's add int 1 to the user rating and print the result. Remember the user input of type string
	// we use the strconv from Go package to add int and string
	//numRating = input + 1 // throw err
	// we use ParseFloat, converts the string to a floating-point number. It which requires two values to be passed in; the string you passed on and the bitsize of it

	/* numRating, err := strconv.ParseFloat(input, 64) // using this will produce the syntax error 'strvconv.ParseFloat: parsing "input\r\n"'.  Because, when we hit input, we ain't only submitting the input(eg 4), there is a trailing character that comes with it, which is \n. we need to find a way to trim the input and \n */
	// we remove the input and use another package called "strings" together with TrimSpace method

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// since Go is alerting that we ain't using numRating and err, everytime you see this error, the special syntax we use is; we move the program to cautiously check if there is error or not

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating:", numRating+1)
	}

}
