package main

import "fmt"

func main() {
    const PI = 3.14
    var l, r, s float64
    fmt.Print("L = ")
    fmt.Scanf("%f", &l)
    r = l / (2 * PI)
    s = PI * r * r
    fmt.Printf("R = %.2f\n", r)
    fmt.Printf("S = %.2f\n", s)
}