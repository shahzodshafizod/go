package main

import (
    "fmt"
    "math"
)

func main() {
    var a float64
    fmt.Print("a = ")
    fmt.Scanf("%f", &a)
    v := math.Pow(a, 3)
    s := 6 * a * a
    fmt.Printf("V = %.3f\nS = %.3f\n", v, s)
}