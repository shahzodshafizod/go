package main

import "fmt"

func main() {
    var m, n int
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
    fmt.Println()
    var mul float32
    for col := 0; col < n; col++ {
        mul = 1
        for row := 0; row < m; row++ {
            mul *= matrix[row][col]
        }
        fmt.Printf("%.2f\t", mul)
    }
}