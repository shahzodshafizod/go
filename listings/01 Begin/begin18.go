package main

import (
    "fmt"
    "math"
)

func main() {
    var a, b, c float64
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    fmt.Print("C = ")
    fmt.Scan(&c)
    ac := math.Abs(c - a)
    bc := math.Abs(c - b)
    mul := ac * bc
    fmt.Printf("AC * BC = %.2f\n", mul)
}