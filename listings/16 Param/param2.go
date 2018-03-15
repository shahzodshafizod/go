package main

import "fmt"

func MaxNum(a []float32, n int) int {
    num := 0
    for index := 1; index < n; index++ {
        if a[index] > a[num] {
            num = index
        }
    }
    return num + 1
}

func main() {
    var n int
    for i := 1; i <= 3; i++ {
        fmt.Printf("N%d = ", i)
        fmt.Scan(&n)
        var array []float32 = make([]float32, n)
        for j := 0; j < n; j++ {
            fmt.Scan(&array[j])
        }
        fmt.Printf("MaxNum = %d\n\n", MaxNum(array, n))
    }
}