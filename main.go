package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println("Hello, Kelechi. Welcome to Golang")
	// DATA TYPES IN GO
	/* string, bool, Integer; (uint8 uint64 int8 int64 uintptr aliases), Floating; (float32, float64),complex
	complex data: array, slices, map, structs, pointers etc
	*/
	// declaring variable
	var user1 string = "Kelechi"
	fmt.Println(user1)
	// alternatively
	user2 := "Duru" // with := you don't need var keyword and data type
	fmt.Println(user2)

	// three main primitive data types on Go; boolean, integars, strings, rune
	// BOOLEAN
	var n bool = true
	fmt.Printf("%v, %T\n", n, n) /* Printf or Sprintf serves as a placeholder.

	Go PLACEHOLDERS
	%v = default format for value
	%d = formatted int
	%f = formatted floating point numbers
	%s = formatted strings
	%t = formatted boolean
	%c = formatting single characters
	%b = formatted base 2(binary)
	%x = formatted base 16(hexadecimal)
	%o = formatted base 8

	EXAMPLE
	name := "Alice"
	age := 25
	balance := 123.45
	isStudent := true

	fmt.Printf("name: %s, age := %d, balance := %f, isStudent := %t\n", name, age, balance, isStudent) */

	// INTEGERS; usigned int only have positive numbers; uint16. signed int can have both the postive or negative numbers; int16.
	m := 42
	fmt.Printf("%v, %T\n", m, m)

	// STRING CONCACTENATION
	s := "welcome to Go programming."
	s2 := "I hope you enjoy the lecture."
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	// RUNE(primitive data); rep with a single qoute ''. whereas string rep UTF8 character, rune rep UTF32 character
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r) // 97, int32. rune is a true type, which is the same thing as talking about int32

	// CONSTANTS(implicit conversion can happen with const but not with var; can interoparate)
	// in naming const, we don't name starting with a cap letter, if we do so, it will export the const to other packages. Instead, we name const as we name other var to keep it scoped to that particular package

	const myConst int = 52 // also written const myConst = 52
	fmt.Printf("%v, %T\n", myConst, myConst)

	// we can add const to a var if they are of the same type
	const myConst1 int = 65
	var myConst2 int = 18
	fmt.Printf("%v, %T\n", myConst1+myConst2, myConst1+myConst2)

	// ARRAYS AND SLICES
	// ARRAYS: size must be indicated inside [], else if becomes a slice. unlimited array rep with three dots [...]
	// DECLARING ARRAY
	// var grades [3]int = [3]int{97, 83, 78}

	grades := [...]int{97, 83, 78} // ... unlimited size of the array
	fmt.Printf("grades: %v", grades)

	var students [3]string
	fmt.Printf("students: %v\n", students)
	students[0] = "Ada"
	students[1] = "Chidi"
	students[2] = "Eze"
	fmt.Printf("students: %v\n", students)

	// FINDING ELEMENT IN THE ARRAY
	fmt.Printf("student #1: %v\n", students[1])

	// FINDING SIZE OF THE ARRAY
	// with len()
	fmt.Printf("Number of students: %v\n", len(students)) // in Go, length is always equal to the declared size of the array even if it contains less elements

	// ARRAY OF ARRAYS
	/* var identifyMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println((identifyMatrix)) */
	// similar code
	var identifyMatrix [3][3]int
	identifyMatrix[0] = [3]int{1, 0, 0}
	identifyMatrix[1] = [3]int{0, 1, 0}
	identifyMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identifyMatrix)

	// SLICES
	// we use the index number eg [2] to add elements in array
	// we use append method eg append(sliceName, sliceElementToAdd)
	var sweetDrink = []string{"mango", "orange", "apple"}
	fmt.Printf("Type of sweetDrink is %T\n", sweetDrink)

	// ADDING ELEMENT IN THE SLICE
	sweetDrink = append(sweetDrink, "cashew", "pear")
	fmt.Println(sweetDrink)

	// SLICING THE SLICE ie making a separate slice from the original slice
	// use [:] eg [2:] = begin slice from index 2 and include all other elements, [2:4] = slice from index 2 but don't include index 4. [:3] == [0:3] = slice from index 0 and exclude index 3.

	/* sweetDrink = append(sweetDrink[2:])
	fmt.Println(sweetDrink)
	sweetDrink = append(sweetDrink[1:4])
	fmt.Println(sweetDrink) */
	sweetDrink = append(sweetDrink[:3])
	fmt.Println(sweetDrink)

	// ADDING AN ELEMENT IN A SLICE
	// interms of memeory mgt, there are 2 ways; the "make keyword" and "new keyword"
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 451
	highScores[2] = 578
	highScores[3] = 118
	//highScores[4] = 736 // without using the append method, 736 will not be added, as the default setting is size of 4

	highScores = append(highScores, 367, 878, 908) // append method will auto reallocate the memory/increase the default size of 4 to 7 to accommodate the new elements, unlike in array which does not, except you use ... in []

	fmt.Println(highScores)

	// sort method in slice
	sort.Ints(highScores)
	fmt.Printf("highScores sorted order: %v\n", highScores)
	// finding boolean from sort method
	//fmt.Println(sort.IntsAreSorted(highScores))
	fmt.Println(slices.IsSorted(highScores))

	// HOW TO REMOVE A VALUE FROM SLICES BASED ON INDEX
	var courses = []string{"reactJs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) // removed "swift" and continued with the other elements
	fmt.Println(courses)

}
