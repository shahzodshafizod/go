package main

import "fmt"

func main() {
    const PI = 3.14
    var r float64
    fmt.Print("R = ")
    fmt.Scanf("%f", &r)
    l := 2 * PI * r
    s := PI * r * r
    fmt.Printf("L = %.3f\nS = %.3f\n", l, s)
}