package main

import (
    "fmt"
    "math"
)

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scan(&n)
    var (
        number float64
        rounded, sum int
    )
    for n > 0 {
        fmt.Scan(&number)
        rounded = int(math.Round(number))
        fmt.Printf("%.d\t", rounded)
        sum += rounded
        n--
    }
    fmt.Printf("\nsum = %d\n", sum)
}