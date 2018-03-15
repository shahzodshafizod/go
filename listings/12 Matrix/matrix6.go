package main

import "fmt"

func main() {
    var (
        m, n int
        d float32
    )
    fmt.Print("M = ")
    fmt.Scan(&m)
    fmt.Print("N = ")
    fmt.Scan(&n)
    fmt.Print("D = ")
    fmt.Scan(&d)
    var array []float32 = make([]float32, n)
    for index, _ := range array {
        fmt.Scan(&array[index])
    }
    var matrix [][]float32 = make([][]float32, m)
    for row, _ := range matrix {
        matrix[row] = make([]float32, n)
        for col, _ := range matrix[row] {
            if row == 0 {
                matrix[row][col] = array[col]
            } else {
                matrix[row][col] = matrix[row-1][col] * d
            }
        }
    }
    fmt.Println()
    for _, array := range matrix {
        for _, value := range array {
            fmt.Printf("%.2f\t", value)
        }
        fmt.Println()
    }
}