package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    k, degree := 0, 1
    for degree < n {
        degree *= 3
        k++
    }
    k--
    fmt.Printf("K= %d\n", k)
}