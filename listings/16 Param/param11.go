package main

import "fmt"

func SortArray(a []float32, n int) {
    // Bubble Sort
    for i := 0; i < n-1; i++ {
        for j := 1; j < n-i; j++ {
            if a[j-1] > a[j] {
                tmp := a[j-1]
                a[j-1] = a[j]
                a[j] = tmp
            }
        }
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
    for i := 1; i <= 3; i++ {
        fmt.Printf("N%d = ", i)
        fmt.Scan(&n)
        var array []float32 = make([]float32, n)
        for index, _ := range array {
            fmt.Scan(&array[index])
        }
        SortArray(array, n)
        printArray(array, n)
    }
}