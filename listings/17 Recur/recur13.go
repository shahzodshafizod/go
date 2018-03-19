package main

import "fmt"

func Palindrome(s []rune, from int, to int) bool {
	if from >= to {
		return true
	}
	return s[from] == s[to] && Palindrome(s, from+1, to-1)
}

func main() {
	var str string
	for i := 0; i < 5; i++ {
		fmt.Scan(&str)
		var codes []rune = []rune(str)
		fmt.Printf("Palindrome:\t%t\n", Palindrome(codes, 0, len(codes)-1))
	}
}
