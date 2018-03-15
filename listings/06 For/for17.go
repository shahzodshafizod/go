package main

import "fmt"

func main() {
    var (
        a float64
        n int
        sum float64
        degree float64 = 1
    )
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 0; i <= n; i++ {
        sum += degree
        degree *= a;
    }
    fmt.Printf("result = %.2f\n", sum)
}