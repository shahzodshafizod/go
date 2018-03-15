package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    var result, number float64 = 0, 1.1
    alomat := 1
    for i := 0; i < n; i++ {
        result += float64(alomat) * number
        number += 0.1
        alomat *= -1
    }
    fmt.Printf("result = %.2f\n", result)
}