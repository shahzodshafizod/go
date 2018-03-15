package main

import "fmt"

func main() {
    var (
        n int
        a, b, s, minS float32
        inited bool
    )
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        fmt.Printf("a%d = ", i)
        fmt.Scan(&a)
        fmt.Printf("b%d = ", i)
        fmt.Scan(&b)
        s = a * b
        if !inited {
            minS, inited = s, true
        } else if s < minS {
            minS = s
        }
    }
    fmt.Printf("minS = %.2f\n", minS)
}