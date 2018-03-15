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
    for row := 0; row < m; row++ {
        sum += matrix[row][k-1]
        mul *= matrix[row][k-1]
    }
    fmt.Printf("sum = %.2f\t\tmultiplication = %.2f\n", sum, mul)
}