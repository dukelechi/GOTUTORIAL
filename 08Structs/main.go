package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang") // alt to class in JS
	// no inheritance in golang; no super or parent
	// very important to use caps in defining struct variables as they will be exported(public)
	// starts with "type keyword"

	kelechi := User{"Kelechi", "kelechi@go.dev", true, 16}
	fmt.Println(kelechi)
	fmt.Printf("kelechi details are: %+v\n", kelechi) // use %+v if it's a structure; gives more details.
	fmt.Printf("Name is %v and email is %v.", kelechi.Name, kelechi.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
