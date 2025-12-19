package main

import "fmt"

// Hello world
// Example for print.
// The main function runs when our program starts
func main() {
	myPrint("hello, world")
	myPrint("goodbye, world")
}

// a function that prints a message
func myPrint(str string) {
	fmt.Println(str)
}

// Add a comment for git fetch
