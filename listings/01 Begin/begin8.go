package main

import "fmt"

func main() {
    var a, b float64
    fmt.Print("a = ")
    fmt.Scan(&a)
    fmt.Print("b = ")
    fmt.Scan(&b)
    aMean := (a + b) / 2
    fmt.Printf("aMean = %.2f\n", aMean)
}