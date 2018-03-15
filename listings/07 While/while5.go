package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    var degree, k = 1, 0
    for degree < n {
        k++
        degree *= 2
    }
    fmt.Printf("K = %d\n", k)
}