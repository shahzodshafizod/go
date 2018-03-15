package main

import "fmt"

func main() {
    const PI = 3.14
    var d float32
    fmt.Print("d = ")
    fmt.Scan(&d)
    l := PI * d
    fmt.Printf("L = %.2f\n", l)
}