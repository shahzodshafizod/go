package main

import (
    "fmt"
    "math"
)

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scan(&n)
    var number, sum, butun float64
    for n > 0 {
        fmt.Scan(&number)
        butun = math.Floor(number)
        sum += butun
        fmt.Printf("%.2f\t", butun)
        n--
    }
    fmt.Printf("\nsum = %.2f\n", sum)
}