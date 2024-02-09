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
	fmt.Printf("Name is %v and email is %v.\n", kelechi.Name, kelechi.Email)
	kelechi.GetUsername() // GetUsername is a method acting on the struct; kelechi
	kelechi.GetStatus()
	kelechi.NewEmail() // this method passes a copy of the new value(in this case; test@go.dev) created in the property of struct kelechi instead of the original value(in this case; kelechi@go.dev). This is the reason why pointers exist in Go. To pass the original value, we use a pointer/reference.
	fmt.Printf("Name is %v and email is %v.\n", kelechi.Name, kelechi.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// METHODS IN GO
// remember, no inheritance or parent in Go. My description below is for understanding purposes
// this is a function that works on the class of an object/parent; works on the property of an object to manipulate its value
// methods in Go are functions that are associtated with a specific type. They enable you to define behaviours or actions that can be performed by INSTANCES of the TYPE.
// method work on the struct type
// the receiver type can be either a VALUE or POINTER RECEIVER. A value receiver operates on a copy of the instance(property/key), while a pointer receiver allows the method to modify the original instance(struct)
// since we don't have parents or classes in Go, we bring our functions into struct, ans thus call them methods

func (u User) GetUsername() { // u is used as var name to rep an instance of the "User" type
	fmt.Println("My username is:", u.Name)
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

// manipulating the email in struct kelechi
func (u User) NewEmail() {
	u.Email = "test@go.dev" // creating a new email; mainpulated the original email
	fmt.Println("The new email address is:", u.Email)
}
