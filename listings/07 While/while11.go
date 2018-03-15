package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    k, sum := 0, 0
    for sum < n {
        k++
        sum += k
    }
    fmt.Printf("K = %d\t\tsum = %d\n", k, sum)
}