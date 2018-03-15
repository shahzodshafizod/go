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
    var row, modifier int
    for col := 0; col < n; col++ {
        if col % 2 == 0 {
            row, modifier = 0, 1
        } else {
            row, modifier = m - 1, -1
        }
        for {
            if row < 0 || row >= m { break }
            fmt.Printf("%.2f\t", matrix[row][col])
            row += modifier
        }
        fmt.Println()
    }
}