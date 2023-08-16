package main

import (
	"fmt"
	"runtime"
)

func main() {
	
	go func() {
		for i := 1; i <= 1000; i++ {
			go evenorodd(i)

		}
		runtime.Goexit()
	}()

}
func evenorodd(num int) {
	if num%2 == 0 {
		fmt.Println(num, "even num")
	} else {
		fmt.Println(num, "odd num")
	}
	

}
