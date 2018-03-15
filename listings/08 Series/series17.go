package main

import "fmt"

func main() {
    var (
        b, number float64
        n int
        printed = false
    )
    fmt.Print("B = ")
    fmt.Scan(&b)
    fmt.Print("N = ")
    fmt.Scan(&n)
    for n > 0 {
        fmt.Scan(&number)
        if !printed && b < number {
            fmt.Printf("%.2f\t", b)
            printed = true
        }
        fmt.Printf("%.2f\t", number)
        n--
    }
    if !printed {
        fmt.Printf("%.2f\n", b)
    }
}