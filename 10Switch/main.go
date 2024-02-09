package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Statement in Golang")

	// creating a dice(ludo) game
	rand.Seed(time.Now().UnixNano()) // generating a random num with time
	diceNumber := rand.Intn(6) + 1   // rand.Initn(6) will print a random non-negative num bw range 0-5. num 6 won't be printed as it's the max num. To solve this, we added 1 to the statement

	fmt.Println("value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move to spot 2")
	case 3:
		fmt.Println("you can move to spot 3")
		fallthrough // once we hit case 3, it auto prints case 4 but not case 5 and 6
	case 4:
		fmt.Println("you can move to spot 4")
	case 5:
		fmt.Println("you can move to spot 5")
	case 6:
		fmt.Println("you can move to spot 6 and role the dice again")
	default:
		fmt.Println("What was that!")
	}

}
