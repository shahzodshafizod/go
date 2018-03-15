package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scan(&n)
	var a, b, result float32
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Printf("%.2f", a)
	switch n {
	case 1:
		result = a + b
		fmt.Print(" + ")
	case 2:
		result = a - b
		fmt.Print(" - ")
	case 3:
		result = a * b
		fmt.Print(" * ")
	case 4:
		result = a / b
		fmt.Print(" / ")
	}
	fmt.Printf("%.2f = %.2f\n", b, result)
}
