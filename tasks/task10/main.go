//Take 2 dimentional array 3 X 3 and assign values .
//Get the sume of Each row and each column data
//Write a func to perform those operations
//Input:
//1 2 3
//4 5 6
//7 8 9
//Output:
//Row 1 Sum: 6
//Row 2 Sum: 15
//Row 3 Sum: 24
//Column 1 Sum: 12
//Column 2 Sum: 15
//Column 3 Sum: 18

package main

import "fmt"

func calculateSums(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		rowSum := 0
		for j := 0; j < cols; j++ {
			rowSum += matrix[i][j]
		}
		fmt.Printf("Row %d Sum: %d\n", i+1, rowSum)
	}

	for j := 0; j < cols; j++ {
		colSum := 0
		for i := 0; i < rows; i++ {
			colSum += matrix[i][j]
		}
		fmt.Printf("Column %d Sum: %d\n", j+1, colSum)
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	calculateSums(matrix)
}
