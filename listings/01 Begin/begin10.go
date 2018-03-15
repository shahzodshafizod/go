package main

import "fmt"

func main() {
    var a, b float64
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    sqrA := a * a
    sqrB := b * b
    sum := sqrA + sqrB
    sub := sqrA - sqrB
    mul := sqrA * sqrB
    div := sqrA / sqrB
    fmt.Printf("sum = %.2f\n", sum)
    fmt.Printf("sub = %.2f\n", sub)
    fmt.Printf("mul = %.2f\n", mul)
    fmt.Printf("div = %.2f\n", div)
}