package main

import "fmt"

func main() {
	var num1 uint8 = 38
	var num2 int = int(num1)
	fmt.Println(num2)
	var num3 int = 65001
	var num4 uint8 = uint8(num3)
	fmt.Println(num4)
	var str1 string = "H"
	var num5 int = int(str1[0])
	fmt.Println(num5)
	var num6 int = 65
	var str2 string = string(num6)
	fmt.Println(str2)
	str3 := fmt.Sprint(num6)
	fmt.Println(str3)
	ok := false
	str4 := fmt.Sprint(ok, num6, "Hello") // no space is required if it is not string
	fmt.Println(str4)
}
