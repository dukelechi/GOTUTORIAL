package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files on Golang")
	// assignment; develop how we can use files(like a pdf, word doc etc); we have some content we want to put on/write in a text file, how can we handle this?
	// we need os and ioutils(inputOutputUtilities) packages
	// also need defer; once we are done reading or writing, the file should be closed

	// STEP 1: WHAT IS THE CONTENT WE WANT TO PUT IN A TEXT FILE?
	content := "Deforestation is one of the major reasons for increasing global warming in the world today. This has caused"

	// STEP 2: WHERE DO WE WANT TO CREATE THIS FILE?
	// in the current directory(folder); 15workingWithFiles
	// "./" means in the current directory, then the title of the file should follow immediately after that, plus the file type, which is ".txt"

	// STEP 2.1
	// os.Create("./deforestation.txt")

	// STEP 2.2; we anticipate error ourselves; file is created but we are ready to anticipate error ourselves
	file, err := os.Create("./deforestation.txt")

	// checking possible error
	if err != nil {
		panic(err) // panic stops/shutdown the execution of the program and show you the error you are facing
	}

	// STEP 3: we write something into this file(deforestation.txt)
	// io.WriteString(file, content)
	//STEP 3.1 if the file executes without and error, we get the length of the text as an int. We need to convert the length into strings and also anticipate for error(a must).
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length of the file:", length)

	// STEP 4: close file; always recomm to use "defer", incase we want to write further code down the line
	defer file.Close()

	// WE WANT TO READ THE FILE AS WELL, LET'S CREATE A SEPERATE METHOD FOR IT, AS SHOWN BELOW THIS PRESENT BLOCK, Main()
	readFile("./deforestation.txt")

}

// REMEMBER: CREATION IS os operation; os package is for creating the file while ioutil package is for writing, reading etc
func readFile(filename string) {
	// ioutil.ReadFile(filename)
	// note: whenever you read the file, it's never being read in the string format but in byte format. Below, the type of data in databyte is byte
	databyte, err := ioutil.ReadFile(filename)
	/* if err != nil {
		panic(err)
	} */
	checkNilErr(err)
	// fmt.Println("Text data inside the file is \n", databyte)
	// to convert the data type, in this case "databyte" into a readable format, we convert to string
	fmt.Println("Text inside the file is \n", string(databyte))
}

// to avoid writing the error check too many times, we can create a function for it and use it
// common func name used is checkNilErr
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
