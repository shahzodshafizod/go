package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scan(&n)
    a := make([]float32, n)
    for index, _ := range a {
        fmt.Scan(&a[index])
    }
    for index := 0; index < n; index += 2 {
        fmt.Printf("%.2f\t", a[index])
    }
    index := len(a) - 1
    if n % 2 != 0 {
        index--
    }
    for ; index >= 0; index -= 2 {
        fmt.Printf("%.2f\t", a[index])
    }
}