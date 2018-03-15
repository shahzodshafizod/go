package main

import "fmt"

func RemoveForInc(a *[]float32, n *int) {
    for index := 1; index < *n; {
        if (*a)[index] < (*a)[index-1] {
            *a = append((*a)[:index], (*a)[index+1:] ...)
            (*n)--
        } else {
            index++
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
        RemoveForInc(&array, &n);
        printArray(array, n)
    }
}