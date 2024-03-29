package main

import "fmt"

func main() {
	slice1 := []int{10, 20, 30}
	arr1 := [5]int{10, 11, 12, 13, 14}
	slice2 := arr1[:]
	slice3 := arr1[:3]
	slice4 := arr1[3:]
	slice5 := arr1[1:4]
	fmt.Println(slice1, slice2, slice3, slice4, slice5, arr1)
	//slice6 := make([]int, 0) //0 does not mean slice is nil
	slice6 := make([]int, 2, 4)
	fmt.Println(slice6)
	slice6[0] = 101
	slice6[1] = 102
	slice6 = append(slice6, 103)
	slice6 = append(slice6, 104)
	fmt.Println(slice6, len(slice6), cap(slice6), &slice6[0])
	slice6 = append(slice6, 105)
	fmt.Println(slice6, len(slice6), cap(slice6), &slice6[0])
	slice7 := make([]int, 2)
	slice7[0] = 87
	slice7[1] = 92
	slice8 := slice7
	slice8[0] = 97
	slice8[1] = 102
	fmt.Println(slice7, slice8)
	slice7 = append(slice7, 103)
	fmt.Println(slice7, slice8)
	slice8[0] = 107
	slice8[1] = 112
	fmt.Println(slice7, slice8)
	//	slice9 := slice8 //shallow copy
	slice10 := make([]int, len(slice8))
	copy(slice10, slice8)
	slice10[0] = 207
	fmt.Println(slice10, slice8)

}
