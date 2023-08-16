//Use fmt package functions like Print,Println,Printf and identify the differences

package main

import "fmt"

func main() {
	//using fmt.print
	fmt.Print("Hello", "world") // Output: Hello world
	//using fmt.println
	fmt.Println("Hello", "world") // Output: Hello world (with a newline)
	//using fmt.Printf
	name := "Megha"
	age := 21
	fmt.Printf("Name:%s, Age:%d\n", name, age) // Output: Name: megha Age: 21 (with a newline)
}
