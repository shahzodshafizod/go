package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    sum := 0
    for i := 0; i <= n; i++ {
        sum += (n + i) * (n + i)
    }
    fmt.Printf("result = %d\n", sum)
}