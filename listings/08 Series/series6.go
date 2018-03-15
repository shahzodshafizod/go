package main

import (
    "fmt"
    "math"
)

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scan(&n)
    var number, kasri float64
    var mul float64 = 1
    for n > 0 {
        fmt.Scan(&number)
        kasri = number - math.Floor(number)
        fmt.Printf("%.2f\t", kasri)
        mul *= kasri
        n--
    }
    fmt.Printf("\nmultiplication = %.6f\n", mul)
}