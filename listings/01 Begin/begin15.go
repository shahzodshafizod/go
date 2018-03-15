package main

import (
    "fmt"
    "math"
)

func main() {
    const PI = 3.14
    var s, d, l float64
    fmt.Print("S = ")
    fmt.Scanf("%f", &s)
    d = math.Sqrt(4 * s / PI)
    l = PI * d
    fmt.Printf("D = %.2f\n", d)
    fmt.Printf("L = %.2f\n", l)
}