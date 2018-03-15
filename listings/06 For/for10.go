package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    var sum float64 = 0
    for i := 1; i <= n; i++ {
        sum += 1 / float64(i)
    }
    fmt.Printf("sum = %.6f\n", sum)
}