package main

import (
    "fmt"
    "math"
)

func main() {
    var a, b float64
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    absA := math.Abs(a)
    absB := math.Abs(b)
    sum := absA + absB
    sub := absA - absB
    mul := absA * absB
    div := absA / absB
    fmt.Printf("sum = %.2f\n", sum)
    fmt.Printf("sub = %.2f\n", sub)
    fmt.Printf("mul = %.2f\n", mul)
    fmt.Printf("div = %.2f\n", div)
}