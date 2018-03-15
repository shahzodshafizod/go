package main

import "fmt"

func Fact2(n int) float64 {
    if n < 2 {
        return 1
    }
    return float64(n) * Fact2(n - 2)
}

func main() {
    var n int
    for i := 1; i <= 5; i++ {
        fmt.Printf("N%d = ", i)
        fmt.Scan(&n)
        fmt.Printf("%d! = %.1f\n", i, Fact2(n))
    }
}