package main

import "fmt"

func main() {
    var n, k, l int
    fmt.Print("N = ")
    fmt.Scan(&n)
    array := make([]float32, n)
    for index, _ := range array {
        fmt.Scan(&array[index])
    }
    fmt.Print("K = ")
    fmt.Scan(&k)
    fmt.Print("L = ")
    fmt.Scan(&l)
    var sum float32
    for index := k-1; index < l; index++ {
        sum += array[index]
    }
    fmt.Printf("sum = %.2f\n", sum)
}