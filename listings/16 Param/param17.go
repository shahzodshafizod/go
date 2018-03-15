package main

import "fmt"

func ArrayToMatrCol(a []float32, k int, m int, n int, b *[][]float32) {
    *b = make([][]float32, m)
    index := 0
    for row, _ := range *b {
        (*b)[row] = make([]float32, n)
    }
    for col := 0; col < n; col++ {
        for row := 0; row < m; row++ {
            if (index < k) {
                (*b)[row][col] = a[index]
                index++
            } else {
                (*b)[row][col] = 0
            }
        }
    }
}

func main() {
    var k, m, n int
    fmt.Print("K = ")
    fmt.Scan(&k)
    var array []float32 = make([]float32, k)
    for index, _ := range array {
        fmt.Scan(&array[index])
    }
    fmt.Print("M = ")
    fmt.Scan(&m)
    fmt.Print("N = ")
    fmt.Scan(&n)
    var matrix [][]float32
    ArrayToMatrCol(array, k, m, n, &matrix)
    for row, _ := range matrix {
        for col, _ := range matrix[row] {
            fmt.Printf("%.2f\t", matrix[row][col])
        }
        fmt.Println()
    }
}