package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	chs := []chan int{ch1, ch2, ch3}

	ch4 := VariadicChan(chs...)

	go func() {
		ch1 <- 10
		ch2 <- 20
		ch3 <- 30
	}()
	for sum1 := range ch4 {
		fmt.Println(sum1)

	}

}

func VariadicChan(chs ...chan int) chan int {
	ch := make(chan int)
	go func() {
		sum := 0
		for _, v := range chs {
			sum = sum + <-v
		}
		ch <- sum
		close(ch)
	}()
	return ch
}
