//Create global variables and local variables with the same name in different scopes
//Identify which variable values are taken within or outside of scopes

package main

import "fmt"

// This is a global variable
var x int = 10

func main() {
	// This is a local variable with the same name as the global variable
	var x int = 5

	fmt.Println("Local x:", x) // Prints the local variable

	// Calling a function that uses the global variable
	printGlobalX()
}

func printGlobalX() {
	// Accessing the global variable
	fmt.Println("Global x:", x)
}
