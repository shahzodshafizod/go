package main

import "fmt"

func Smooth1(a []float32, n int) {
    var sum float32
    for index, _ := range a {
        sum += a[index]
        a[index] = sum / float32(index + 1)
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
        Smooth1(array, n)
        printArray(array, n)
    }
}