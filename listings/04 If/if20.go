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
    ab := math.Abs(b - a)
    ac := math.Abs(c - a)
    if ab < ac {
        fmt.Printf("closest point: %.2f\nclosest distance: %.2f\n", b, ab)
    } else {
        fmt.Printf("closest point: %.2f\nclosest distance: %.2f\n", c, ac)
    }
}