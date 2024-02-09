package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// formating the time
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // format is fixed from the time pkg

	// creating your own date
	createdDate := time.Date(2024, time.July, 10, 21, 31, 24, 0, time.UTC)
	fmt.Println(createdDate)

	// formating createdDate to look beautiful
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
