package main

import "fmt"

func MaxElem(a []int, n int) int {
    if n <= 0 {
        return 0
    }
    maximal := a[n-1]
    if n > 0 {
        data := MaxElem(a, n-1)
        if data > maximal {
            maximal = data
        }
    }
    return maximal
}

func main() {
    var n int
    for i := 1; i <= 3; i++ {
        fmt.Printf("N%d = ", i)
        fmt.Scan(&n)
        var array []int = make([]int, n)
        for index, _ := range array {
            fmt.Scan(&array[index])
        }
        maximal := MaxElem(array, n)
        fmt.Printf("maximal = %d\n\n", maximal)
    }
}