package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	/* for d := 0; d < len(days); i++ {
		fmt.Println(days[d])
	} */

	/* for i := range days {
		fmt.Println(days[i])
	} */

	/* for index, day := range days {
		fmt.Printf("index is %v and value is %T\n", index, day)
	} */

	/* for _, day := range days {
		fmt.Printf("index is and value is %T\n", day)
	} */

	// BREAK AND CONTINUE
	rogueValue := 1

	for rogueValue < 10 {
		fmt.Println("value is:", rogueValue)
		rogueValue++

		/* if rogueValue == 5 {
			break // loop stops and does continue iteration
		} */

		if rogueValue == 5 {
			rogueValue++
			continue // value 5 is excluded from the loop and continued to the next value. rogueValue++ makes the increment after 5 possible.
		}

		if rogueValue == 2 {
			goto lco
		}

	}

	// GOTO STATEMENT
	// use can use any label for this operation. In the example below, we used lco

lco:
	fmt.Println("jumping at learncodeonline.in") // this will be executed, due to line 43-45

}
