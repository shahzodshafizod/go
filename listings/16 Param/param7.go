package main

import "fmt"

func Smooth3(a []float32, n int) {
    if n == 1 { return }
    var prev, curr float32
    for index, _ := range a {
        curr = a[index]
        if index == 0 {
            a[index] = (curr + a[index+1]) / 2.0
        } else if index == n - 1 {
            a[index] = (prev + curr) / 2.0
        } else {
            a[index] = (prev + curr + a[index+1]) / 3.0
        }
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
        Smooth3(array, n)
        printArray(array, n)
    }
}