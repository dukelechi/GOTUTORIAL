package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	// ASSIGNMENT: we want to take input from the user, store and return a value to him
	// to achieve this, we need a GO libray(from go.dev; we use package bufio(can read I/O, ie Input/Output) and OS ) and the "comma, ok" syntax
	// in Go, we packages are inplicitly imported
	fmt.Println(welcome)

	// first, create a reader from bufio, ie reader := bufio
	// create a new reader from the library, OS(from where we read), ie reader := bufio.NewReader(OS)
	// create a standard reader from Input/Output, ie reader := bufio.NewReader(OS.Stdin)
	// any method starting in cap means it's public, eg .NewReader, .Stdin.
	// enter message the user will read and input value

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:") // no need to include \n as fmt.Println also has it. we use it for fmt.Printf

	// what ever the reader reads, we want to store it in a variable. This is where "comma OK || err OR" syntax comes into play

	// comma OK || err OR; similar to Try... Catch in javaScript
	// in Go, we don't have "try catch", if something goes wrong, there is no line to catch that. Go expects you to treat errors like a variable; true or false or another value. This is where comma OK || err OR comes in
	// comma OK || comma OR means that you can either get an input or receive an error which can be indicated with "err" of if we don't care about the error, we can indicate it with an underscore "_", if you don't care about the input, we can also use the underscore

	// contd from assignment steps
	// we use key word "input", err or not, then reader, we are interested in read string until how long indicated with \n
	// uses of \n in Go; escape sequence for a newline; in other contexts to indicate the end of a line/paragraph. Also used in "fmt package" for formating strings, and in "bufio package" for reading and writing data to files

	input, _ := reader.ReadString('\n') // standard writing; input, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of this rating is %T", input)

}
