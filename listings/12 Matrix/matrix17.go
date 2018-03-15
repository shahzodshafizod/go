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
    var sum, mul float32 = 0, 1
    for col := 0; col < n; col++ {
        sum += matrix[k-1][col]
        mul *= matrix[k-1][col]
    }
    fmt.Printf("sum = %.2f\t\tmultiplication = %.2f\n", sum, mul)
}