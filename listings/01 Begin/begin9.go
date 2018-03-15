package main

import (
    "fmt"
    "math"
)

func main() {
    var a, b float64
    fmt.Print("a [positive] = ")
    fmt.Scan(&a)
    fmt.Print("b [positive] = ")
    fmt.Scan(&b)
    gMean := math.Sqrt(a * b)
    fmt.Printf("gMean = %.2f\n", gMean)
}