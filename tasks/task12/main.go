//Find the second biggest number in an array/slice

//Find the second smallest number in an array/slice

//Write everything using functions.

//SecondBiggest(arr []int)int

//SecondSmallest(arr []int)int

package main

import (
	"fmt"
	"sort"
)

func SecondBiggest(arr []int) int {
	if len(arr) < 2 {
		fmt.Println("Array should have at least two elements.")
		return -1
	}

	// Sort the array in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	return arr[1]
}

func SecondSmallest(arr []int) int {
	if len(arr) < 2 {
		fmt.Println("Array should have at least two elements.")
		return -1
	}

	// Sort the array in ascending order
	sort.Ints(arr)

	return arr[1]
}

func main() {
	array := []int{5, 8, 3, 1, 9, 4, 7, 2, 6}

	secondBiggest := SecondBiggest(array)
	fmt.Println("Second Biggest:", secondBiggest)

	secondSmallest := SecondSmallest(array)
	fmt.Println("Second Smallest:", secondSmallest)
}
