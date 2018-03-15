package main

import "fmt"

func main() {
    var m, n, k int
    fmt.Print("M = ")
    fmt.Scan(&m)
    fmt.Print("N = ")
    fmt.Scan(&n)
    var matrix [][]float32 = make([][]float32, m)
    for row, _ := range matrix {
        matrix[row] = make([]float32, n)
        for col, _ := range matrix[row] {
            fmt.Scan(&matrix[row][col])
        }
    }
    fmt.Print("K = ")
    fmt.Scan(&k)
    fmt.Println()
    for col := 0; col < n; col++ {
        fmt.Printf("%.2f\t", matrix[k-1][col])
    }
}