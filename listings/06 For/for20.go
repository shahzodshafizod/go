package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    var sum, fact float64 = 0, 1
    for i := 1; i <= n; i++ {
        fact *= float64(i)
        sum += fact
    }
    fmt.Printf("result = %.2f\n", sum)
}