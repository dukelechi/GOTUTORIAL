package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")
	// pointers passes the memory address of variables so it 100% guarantees that whatever you are passing on is the actual value from the memory that is being passed on.
	// pointer is simply the direct reference to the memory address.
	// Ref means & (shows the memory address)
	// pointer is indicated with asterix "*" (shows the actual value)
	// pointer guarantees the operation will be carried out in the actual value and not a copy of it

	/* var ptr *int
	fmt.Println("value of pointer is,", ptr)
	fmt.Printf("type %T", ptr) */

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("value of pointer is", ptr)         // memory address(indicated with "&")
	fmt.Println("value of actual pointer is", *ptr) // 23

	*ptr = *ptr + 2
	fmt.Println("New value is:", myNumber)

}
