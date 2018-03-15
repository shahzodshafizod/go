package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    var fact float64 = 1
    for i := 2; i <= n; i++ {
        fact *= float64(i)
    }
    fmt.Printf("%d! = %.2f\n", n, fact)
}