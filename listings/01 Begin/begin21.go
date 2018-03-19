package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64
	fmt.Print("x1 = ")
	fmt.Scan(&x1)

	fmt.Print("y1 = ")
	fmt.Scan(&y1)

	fmt.Print("x2 = ")
	fmt.Scan(&x2)

	fmt.Print("y2 = ")
	fmt.Scan(&y2)

	fmt.Print("x3 = ")
	fmt.Scan(&x3)

	fmt.Print("y3 = ")
	fmt.Scan(&y3)

	var a float64 = math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	var b float64 = math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	var c float64 = math.Sqrt(math.Pow(x3-x1, 2) + math.Pow(y3-y1, 2))
	var P float64 = a + b + c
	var p float64 = P / 2
	s := math.Sqrt(p * (p - a) * (p - b) * (p - c))
	fmt.Printf("P = %.2f\t\tS = %.2f\n", P, s)
}
