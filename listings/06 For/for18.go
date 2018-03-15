package main

import "fmt"

func main() {
    var (
        a float64
        n int
        result float64
        degree float64 = 1
        alomat int = 1
    )
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 0; i <= n; i++ {
        result += float64(alomat) * degree
        degree *= a
        alomat *= -1
    }
    fmt.Printf("result = %.2f\n", result)
}