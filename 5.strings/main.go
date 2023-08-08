package main

import "fmt"

var (
	num1 int
	num2 float32
	num3 float64
)

func main() {
	var str1 string = "Hello World"
	fmt.Println(str1)
	fmt.Println("lenth of str1", len(str1))
	str1 = "HELLO"
	fmt.Println("length of str1", len(str1))
	num1 := 10000
	fmt.Println(num1)

}
