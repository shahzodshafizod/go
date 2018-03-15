package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    sum := 0
    for i := 1; i <= n; i++ {
        sum += 2 * i - 1
        fmt.Printf("%d\t", sum)
    }
}