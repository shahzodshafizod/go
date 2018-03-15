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
    var col, modifier int
    for row, _ := range matrix {
        if row % 2 == 0 {
            col, modifier = 0, 1
        } else {
            col, modifier = n - 1, -1
        }
        for {
            if col < 0 || col >= n { break }
            fmt.Printf("%.2f\t", matrix[row][col])
            col += modifier
        }
        fmt.Println()
    }
}