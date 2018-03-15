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
    for row := 1; row < m; row += 2 {
        for col, _ := range matrix[row] {
            fmt.Printf("%.2f\t", matrix[row][col])
        }
        fmt.Println()
    }
}