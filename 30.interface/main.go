package main

import "fmt"

func main() {

	var icalc ICalculator
	if icalc == nil {
		fmt.Println("nil")
	}
	it1 := &IntType{A: 300, B: 200}
	str1 := &StringType{A: "hello", B: "Megha"}
	icalc = it1
	r1 := icalc.Add()
	fmt.Println("Addtion of IntType:", r1)
	it1.display()
	icalc = str1

}

type ICalculator interface {
	Add() any
}

type StringType struct {
	A, B   string
	Result string
}
type IntType struct {
	A, B   int
	Result int
}

func (i *IntType) Add() any {
	i.Result = i.A + i.B
	return i.Result
}
func (i *StringType) Add() any {
	i.Result = i.A + i.B
	return i.Result
}
func (i *IntType) display() {
	fmt.Println("A:", i.A, "B:", i.B)
}
func (i *StringType) display() {
	fmt.Println("A:", i.A, "B:", i.B)
}
