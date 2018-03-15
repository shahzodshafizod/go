package main

import "fmt"

func CircleS(r float32) float32 {
    return 3.14 * r * r
}

func main() {
    var r, s float32
    for i := 1; i <= 3; i++ {
        fmt.Printf("R%d = ", i)
        fmt.Scan(&r)
        s = CircleS(r)
        fmt.Printf("S%d = %.2f\n\n", i, s)
    }
}