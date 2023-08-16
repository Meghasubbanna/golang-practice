// create variables of 14 types including rune, byte and also take 2 interface{} type variables;For one assign int and the other one assign float32 values.
// assign values to them and perform addition and multiplication, print the details
package main

import "fmt"

func main() {
	var intVar int = 10
	var int8Var int8 = 20
	var int16Var int16 = 30
	var int32Var int32 = 40
	var int64Var int64 = 50

	var uintVar uint = 60
	var uint8Var uint8 = 70
	var uint16Var uint16 = 80
	var uint32Var uint32 = 90
	var uint64Var uint64 = 100

	var float32Var float32 = 3.14
	var float64Var float64 = 6.28

	var runeVar rune = 'A'
	var byteVar byte = 0x41 // ASCII code for 'A'

	var iface1 any
	iface1 = 10
	var iface2 any
	iface2 = 6.8
	result1 := float64(intVar) + float64(int8Var) + float64(int16Var) + float64(int32Var) + float64(int64Var) + float64(uintVar) + float64(uint8Var) + float64(uint16Var) + float64(uint32Var) + float64(uint64Var) + float64(float32Var) + float64(float64Var) + float64(runeVar) + float64(byteVar) + float64(iface1.(int)) + iface2.(float64)
	fmt.Println(result1)
}
