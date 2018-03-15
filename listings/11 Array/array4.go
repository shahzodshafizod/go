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
    var array [] float32 = make([]float32, n)
    for index, _ := range array {
        if index == 0 {
            array[index] = a
        } else {
            array[index] = array[index-1] * d
        }
    }
    for _, value := range array {
        fmt.Printf("%.2f\t", value)
    }
}