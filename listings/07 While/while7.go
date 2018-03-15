package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    k := 1
    for k * k <= n {
        k++
    }
    fmt.Printf("K = %d\n", k)
}