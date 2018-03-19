package main

import "fmt"

func DigitCount(s string) int {
	var codes []rune = []rune(s)
	var count int = 0
	for _, value := range codes {
		if value >= 48 && value <= 57 {
			count++
		}
	}
	return count
}

func main() {
	var str string
	for i := 0; i < 5; i++ {
		fmt.Scan(&str)
		fmt.Printf("digit count:\t%d\n", DigitCount(str))
	}
}
