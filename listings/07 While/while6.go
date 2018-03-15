package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    var fact2, tmp float64 = 1, float64(n)
    for tmp > 0 {
        fact2 *= tmp
        tmp -= 2
    }
    fmt.Printf("%d!! = %.2f\n", n, fact2)
}