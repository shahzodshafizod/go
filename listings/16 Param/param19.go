package main

import "fmt"

func Norm1(a [][]float32, m int, n int) float32 {
    var max, sum float32
    for col := 0; col < n; col++ {
        sum = 0
        for row := 0; row < m; row++ {
            sum += a[row][col]
        }
        if col == 0 {
            max = sum
        } else if sum > max {
            max = sum
        }
    }
    return max;
}

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
    for row, _ := range matrix {
        fmt.Printf("Norm1(A, %d, %d) = %.2f\n", row+1, n, Norm1(matrix, row+1, n))
    }
}