package main

import "fmt"

func MinElem(a []int, n int) int {
    min := a[0]
    for index := 1; index < n; index++ {
        if a[index] < min {
            min = a[index]
        }
    }
    return min
}

func main() {
    var n int
    for i := 1; i <= 3; i++ {
        fmt.Printf("N%d = ", i)
        fmt.Scan(&n)
        array := make([]int, n)
        for j := 0; j < n; j++ {
            fmt.Scan(&array[j])
        }
        fmt.Printf("Min = %d\n\n", MinElem(array, n))
    }
}