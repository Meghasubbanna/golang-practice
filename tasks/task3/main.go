//Create constants and perform mathematical calc on constants

//Addition --> +

//Substration --> -

//Multiplecation --> *

//Division--> /

//Modulous --> %

package main

import "fmt"

func main() {
	const (
		a = 10
		b = 5
	)

	// Addition
	addResult := a + b
	fmt.Println("Addition:", addResult)

	// Subtraction
	subResult := a - b
	fmt.Println("Subtraction:", subResult)

	// Multiplication
	mulResult := a * b
	fmt.Println("Multiplication:", mulResult)

	// Division
	divResult := a / b
	fmt.Println("Division:", divResult)

	// Modulo
	modResult := a % b
	fmt.Println("Modulo:", modResult)
}
