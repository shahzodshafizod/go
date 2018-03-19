package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scanf("%d", &n)
	seconds := n / 60
	fmt.Printf("seconds = %d\n", seconds)
}
