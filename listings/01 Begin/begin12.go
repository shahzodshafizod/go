package main

import (
    "fmt"
    "math"
)

func main() {
    var a, b, c, p float64
    fmt.Print("a = ")
    fmt.Scan(&a)
    fmt.Print("b = ")
    fmt.Scan(&b)
    c = math.Sqrt(a*a + b*b)
    p = a + b + c
    fmt.Printf("c = %.2f\n", c)
    fmt.Printf("P = %.2f\n", p)
}