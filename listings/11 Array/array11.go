package main

import "fmt"

func main() {
    var n, k int
    fmt.Print("N = ")
    fmt.Scan(&n)
    // var a []float32 = make([]float32, n)
    a := make([]float32, n)
    for index, _ := range a {
        fmt.Scan(&a[index])
    }
    fmt.Print("K = ")
    fmt.Scan(&k)
    for index := k-1; index < n; index += k {
        fmt.Printf("%.2f\t", a[index])
    }
}