package main

import "fmt"

func main() {
    var (
        n, maxIndex int
        m, v, p, maxP float32
        inited bool
    )
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        fmt.Printf("m%d = ", i)
        fmt.Scan(&m)
        fmt.Printf("v%d = ", i)
        fmt.Scan(&v)
        p = m / v
        if !inited {
            maxP, maxIndex, inited = p, i, true
        } else if p > maxP {
            maxP, maxIndex = p, i
        }
    }
    fmt.Printf("index = %d\t\tmaxP = %.2f\n", maxIndex, maxP)
}