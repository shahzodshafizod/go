package main

import "fmt"

func main() {
    var (
        n int
        a, b, p, maxP float32
        inited bool
    )
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        fmt.Printf("a%d = ", i)
        fmt.Scan(&a)
        fmt.Printf("b%d = ", i)
        fmt.Scan(&b)
        p = 2 * (a + b)
        if !inited {
            maxP, inited = p, true
        } else if p > maxP {
            maxP = p
        }
    }
    fmt.Printf("maxP = %.2f\n", maxP)
}