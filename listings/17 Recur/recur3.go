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

func main() {
    var (
        x float64
        n int
    )
    fmt.Print("X = ")
    fmt.Scan(&x)
    for i := 1; i <= 5; i++ {
        fmt.Printf("N%d = ", i)
        fmt.Scan(&n)
        fmt.Printf("PowerN(%.2f, %d) = %E\n", x, n, PowerN(x, n))
    }
}