package main

import "fmt"

const LoginToken string = "sjsksk" //using PascalCase means the login token is visible in public/global scope

func main() {
	var username string = "Kelechi"
	fmt.Printf("variable is a type of: %T\n", username)

	// BOOLEAN
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is a type of: %T\n", isLoggedIn)

	// int
	var smallVal uint8 = 255 // uint8 can only accommodate up to 255. number 256 and above will throw error. this can be solved using "int", which can accommodate more values. Alt, you can use uint16 and above but the all have limits. To find out the limits, use doc in Golang specifications, look out for methods, in the types, we have the numeric types and it gives all the details. Alt, use this shortcut 2 **n. for uint8, 2**8 = 256, which means uint8 can only accomm 0 - 255. for uint16 = 2**16 = 65,536; can accomm 0 - 65,535 etc
	fmt.Println(smallVal)
	fmt.Printf("variable is a type of: %T\n", smallVal)

	// FLOAT
	var smallFloat float32 = 783.778888899888776677 // if we use float64, we have more precision with the decimal places(float)
	fmt.Println(smallFloat)
	fmt.Printf("variable is a type of: %T\n", smallFloat)

	// DEFAULT VALUES AND SOME ALIASES
	var anotherVariable int // declaring a var without initializing it prints out 0
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T\n", anotherVariable)

	// IMPLICIT TYPE; Go auto understands the var type
	var website = "learncodeonline.in"
	fmt.Printf("value: %v, Type: %T\n", website, website)

	//NO VAR STYLE; using := is only allowed inside a method and not outside/ in the global scope
	numberOfUser := 30000
	fmt.Printf("value: %v, Type: %T\n", numberOfUser, numberOfUser)

	// CONSTANTS; see line 5. Also, camelCase scope the const to a local scope
	fmt.Printf("value: %v, Type: %T\n", LoginToken, LoginToken)

}
