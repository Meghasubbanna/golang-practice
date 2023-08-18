package main

import (
	"bufio"
	"fmt"
	"os"
)

// var (
// 	lineBreakRegExp = regexp.MustCompile(`\r?\n?.`)
// )

func main() {
	ch := make(chan string)
	file, err := os.Open("/Users/mrs/Documents/learngo/data/dataa.txt")
	if err != nil {
		fmt.Println("error1")
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	for i, line := range fileLines {
		go countChars(ch, i)
		ch <- line
		go countWords(ch, i)
		ch <- line
		go countVowelsAndConsonants(ch, i)
		ch <- line
	}
}
func countChars(ch chan string, i int) {
	str := <-ch
	fmt.Println("length of line", i, ":", len(str))
}

func countWords(ch chan string, i int) {
	str := <-ch
	count := 0
	for _, val := range str {
		if val == ' ' {
			count++
		}
	}
	count++
	fmt.Println("Number of words in line", i, ":", count)
}
func countVowelsAndConsonants(ch chan string, i int) {
	str := <-ch
	var vow, cons, spec = 0, 0, 0
	for _, val := range str {
		if val == 65 || val == 69 || val == 73 || val == 79 || val == 85 || val == 97 || val == 101 || val == 105 || val == 111 || val == 117 {
			vow++
		} else if val >= 65 && val <= 91 || val <= 97 && val <= 121 {
			cons++
		} else {
			spec++
		}
	}
	fmt.Println("Number of vowles in line", i, ":", vow, "\n number of consonents", i, ":", cons, "\n number of special chars in line", i, ":", spec)
}
