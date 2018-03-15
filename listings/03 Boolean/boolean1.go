package main

import "fmt"

func main() {
	var a int
	fmt.Print("A = ")
	fmt.Scanf("%d", &a)
	isPositive := a > 0
	fmt.Printf("Positive = %t\n", isPositive)
}