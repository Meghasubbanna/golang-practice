// Create variables of all number types . 14 datatypes including rune and byte
// declare and assign values of all variables
// perform addition and multiplecation, assign the result to a another variable, print it.
package main

import "fmt"

func main() {
	// Integer types
	var intVar int = 10
	var int8Var int8 = 20
	var int16Var int16 = 30
	var int32Var int32 = 40
	//var int64Var int64 = 50

	// Unsigned integer types
	var uintVar uint = 60
	var uint8Var uint8 = 70
	var uint16Var uint16 = 80
	var uint32Var uint32 = 90
	//var uint64Var uint64 = 100

	// Floating-point types
	var float32Var float32 = 3.14
	var float64Var float64 = 6.28

	//var runeVar rune = 'A'  //rune is alias for int32 and it is used to represent unicode points
	//var byteVar byte = 0x41 // ASCII code for 'A'(65 decimal).byte is an alias for uint8.

	// Performing addition and multiplication
	intAddResult := intVar + int(int8Var)
	intMulResult := int32Var * int32(int16Var)
	uintAddResult := uintVar + uint(uint8Var)
	uintMulResult := uint32Var * uint32(uint16Var)
	floatAddResult := float32(float64(float32Var) + float64Var)
	floatMulResult := float32(float64(float32Var) * float64Var)

	// Printing the results
	fmt.Println("Integer Addition result:", intAddResult)
	fmt.Println("Integer Multiplication result:", intMulResult)
	fmt.Println("Unsigned Integer Addition result:", uintAddResult)
	fmt.Println("Unsigned Integer Multiplication result:", uintMulResult)
	fmt.Println("Floating-Point Addition result:", floatAddResult)
	fmt.Println("Floating-Point Multiplication result:", floatMulResult)
}
