package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string) // map can contain any type of data. Size is not defined

	languages["JS"] = "javascript"
	languages["PY"] = "python"
	languages["RB"] = "ruby"

	fmt.Println("list of all languages:", languages)
	fmt.Println("JS short for:", languages["JS"])
	fmt.Println("RB short for:", languages["RB"])

	// DELETE A KEY
	delete(languages, "RB")
	fmt.Println("list of all languages:", languages)

	// LOOPS ARE INTERESTING IN GOLANG
	// basic syntax
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	/* for _, value := range languages {
	fmt.Printf("For key v, value is %v\n", value) */

}
