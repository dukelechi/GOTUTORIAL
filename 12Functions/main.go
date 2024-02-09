package main

import "fmt"

// DEFAULT FUNCTION; entry point for go programming
func main() {
	fmt.Println("Welcome to functions in Golang")
	greeter()

	// method to add 2 number; first, define the func/method
	result := adder(3, 5)
	fmt.Println("Result is:", result)

	/* proResult := proAdder(4, 6, 8, 9, 10, 7)
	fmt.Println("Pro result is:", proResult) */

	proResult, proAdderMessage := proAdder(4, 6, 8, 9, 10, 7)
	fmt.Printf("Pro result is %v and message result is %v\n", proResult, proAdderMessage)

}

// CREATING OWN FUNCTION
func greeter() {
	fmt.Println("Namstey from golang") // namstey; indian word for hello
} // this function call, even though we used the print func will not be displayed/run until we call it in func main(); default, entry point of Golang
// any func written inside func main() will not run. Go only permits you to call the func inside the entry/default func

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// assignment; write a func that adds unlimited numbers/value
/* func proAdder(values ...int) int { // std synta; func proAdder(values... int), creates a slice of unkown size
	total := 0
	for _, val := range values { // also be written as for index, val := range values. see loops for expl
		total += val
	}
	return total

} */

func proAdder(values ...int) (int, string) { // we want to also return a message, alongside the value, so we specify the data type to be returned, and since we are returning two values, we enclose them in braces; ie (int, string)
	total := 0
	for _, val := range values { // also be written as for index, val := range values. see loops for expl
		total += val
	}
	return total, "Hi, pro result function"
}
