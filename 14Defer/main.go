package main

import "fmt"

func main() {
	// DEFER FUNCTION; delays a function call and follows a LIFO sequence(Last In First Out)

	/* fmt.Println("Hello")
	defer fmt.Println("Welcome to Backend Development") // will be returned last
	fmt.Println("Kelechi") */

	defer fmt.Println("World") // 5th ie Last Out
	defer fmt.Println("One")   // 4th
	defer fmt.Println("Two")   // 3rd
	fmt.Println("Hello")       // 1st print
	myDefer()                  // 2nd print because after "Hello", the next line it will see is myDefer()

	// world, One, Two
	// 0,1,2,3,4
	// Hello, 43210, Two, One, World

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i) // will follow LIFO; 4,3,2,1,0 when printing
	}
}
