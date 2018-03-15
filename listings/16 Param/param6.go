package main

import "fmt"

func Smooth2(a []float32, n int) {
    var prev, curr float32
    prev = a[0]
    for index := 1; index < n; index++ {
        curr = a[index]
        a[index] = (prev + curr) / 2.0
        prev = curr
    }
}

func printArray(a []float32, n int) {
    for index, _ := range a {
        fmt.Printf("%.2f\t", a[index])
    }
    fmt.Println()
}

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scan(&n)
    var array []float32 = make([]float32, n)
    for index, _ := range array {
        fmt.Scan(&array[index])
    }
    fmt.Println()
    for i := 0; i < 5; i++ {
        Smooth2(array, n)
        printArray(array, n)
    }
}