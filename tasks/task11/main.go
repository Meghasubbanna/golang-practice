//Create and assign values to 2d array 3X3 array
//store rows to columns and columns to rows in the array
//Write a func to perform those operations
//Input:
//1 2 3
//4 5 6
//7 8 9
//Output:
//1 4 7
//2 5 8
//3 6 9

package main

import "fmt"

func transposeArray(arr [3][3]int) [3][3]int {
	var transposed [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			transposed[i][j] = arr[j][i]
		}
	}

	return transposed
}

func main() {
	arr := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	transposed := transposeArray(arr)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\t", transposed[i][j])
		}
		fmt.Println()
	}
}
