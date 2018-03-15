package main

import "fmt"

func main() {
    var (
        a float64
        n int
        degree float64 = 1
    )
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        degree *= a;
    }
    fmt.Printf("degree = %.2f\n", degree)
}