package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    var mul, number float64 = 1, 1
    for i := 0; i < n; i++ {
        number += 0.1
        mul *= number
    }
    fmt.Printf("result = %.6f\n", mul)
}