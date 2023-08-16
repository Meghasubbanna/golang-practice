package main

import "fmt"

func main() {

	// single line commenting

	/*
		Multiline
		Commenting
	*/
	str1 := `"Hello World"`
	println(str1)

	a, b := 10, 20
	t := b
	b = a
	a = t

	a, b = b, a // swapping in go

	c := 30
	fmt.Println("Before swap", a, b, c)

	a, b, c = c, a, b
	fmt.Println("After swap", a, b, c)

}
