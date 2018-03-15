package main

import "fmt"

func MinmaxNum(a []float32, n int) (int, int) {
    minIndex, maxIndex := 0, 0
    for index := 1; index < n; index++ {
        if a[index] > a[maxIndex] {
            maxIndex = index
        }
        if a[index] < a[minIndex] {
            minIndex = index
        }
    }
    return minIndex+1, maxIndex+1
}

func main() {
    var n, maxNum, minNum int
    for i := 1; i <= 3; i++ {
        fmt.Printf("N%d = ", i)
        fmt.Scan(&n)
        var array []float32 = make([]float32, n)
        for j := 0; j < n; j++ {
            fmt.Scan(&array[j])
        }
        minNum, maxNum = MinmaxNum(array, n)
        fmt.Printf("minNum = %d\t\tmaxNum = %d\n\n", minNum, maxNum)
    }
}