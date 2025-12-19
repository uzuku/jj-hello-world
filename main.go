package main

import "fmt"

// Hello world
// Example for print.
func main() {
	myPrint("hello, world")
	myPrint("goodbye, world")
}

// a function that prints a message
func myPrint(str string) {
	fmt.Println(str)
}
