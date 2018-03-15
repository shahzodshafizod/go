package main

import "fmt"

func PowerN(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    if n < 0 {
        return 1 / PowerN(x, -n)
    }
    if n % 2 == 0 {
        half := PowerN(x, n/2)
        return half * half
    }
    return x * PowerN(x, n-1)
}

func RootK(x float64, k int, n int) float64 {
    if n == 0 {
        return 1
    }
    prev := RootK(x, k, n-1)
    return prev - (prev - x / PowerN(prev, k-1)) / float64(k)
}

func main() {
    var (
        x, root float64
        k, n int
    )
    fmt.Print("X = ")
    fmt.Scan(&x)
    fmt.Print("K = ")
    fmt.Scan(&k)
    for i := 1; i <= 6; i++ {
        fmt.Printf("N%d = ", i)
        fmt.Scan(&n)
        root = RootK(x, k, n)
        fmt.Printf("%.8f\n\n", root)
    }
}