package main

import "fmt"

func Hill(a []float32, n int) {
    if n == 1 { return }
    from, to := 0, n-1
    for iter := n/2; iter > 0; iter-- {
        for index := to; index > from; index-- {
            if a[index-1] > a[index] {
                tmp := a[index-1]
                a[index-1] = a[index]
                a[index] = tmp
            }
        }
        from++
        for index := from; index < to; index++ {
            if a[index] < a[index+1] {
                tmp := a[index+1]
                a[index+1] = a[index]
                a[index] = tmp
            }
        }
        to--
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
        Hill(array, n)
        printArray(array, n)
    }
}