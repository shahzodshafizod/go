package main

import (
    "fmt"
    "math"
)

func TriangleP(a float64, h float64) float64 {
    b := math.Sqrt(math.Pow(a/2, 2) + h*h)
    return a + 2*b
}

func main() {
    var a, h, p float64
    for i := 1; i <= 3; i++ {
        fmt.Printf("a%d = ", i)
        fmt.Scan(&a)
        fmt.Printf("h%d = ", i)
        fmt.Scan(&h)
        p = TriangleP(a, h)
        fmt.Printf("P%d = %.2f\n\n", i, p)
    }
}