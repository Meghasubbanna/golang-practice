//In a given array find the biggest and smallest number in an array.
//Write a function getBigAndSmall(arr [10]int)(int,int)

package main

import "fmt"

func getBigAndSmall(arr [10]int) (int, int) {
	biggest := arr[0]
	smallest := arr[0]

	for _, num := range arr {
		if num > biggest {
			biggest = num
		}
		if num < smallest {
			smallest = num
		}
	}

	return biggest, smallest
}

func main() {
	arr := [10]int{23, 45, 12, 67, 9, 53, 32, 8, 91, 17}

	biggest, smallest := getBigAndSmall(arr)

	fmt.Printf("Biggest number: %d\n", biggest)
	fmt.Printf("Smallest number: %d\n", smallest)
}
