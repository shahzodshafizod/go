package main

import "fmt"

func main() {
    var (
        n int
        a, d float32
    )
    fmt.Print("N = ")
    fmt.Scan(&n)
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("D = ")
    fmt.Scan(&d)
    var arr []float32 = make([]float32, n)
    for index, _ := range arr {
        if index == 0 {
            arr[index] = a
        } else {
            arr[index] = arr[index-1] + d
        }
    }
    for _, value := range arr {
        fmt.Printf("%.2f\t", value)
    }
}