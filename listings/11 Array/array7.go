package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scan(&n)
    var array []float32 = make([]float32, n)
    for index, _ := range array {
        fmt.Scan(&array[index])
    }
    for index := len(array)-1; index >= 0; index-- {
        fmt.Printf("%.2f\t", array[index])
    }
}