package main

import "fmt"

func main() {
    var m, n int
    fmt.Print("M = ")
    fmt.Scan(&m)
    fmt.Print("N = ")
    fmt.Scan(&n)
    var array []float32 = make([]float32, n)
    for index, _ := range array {
        fmt.Scan(&array[index])
    }
    var matrix [][]float32 = make([][]float32, m)
    for row, _ := range matrix {
        matrix[row] = make([]float32, n)
        for col, _ := range matrix[row] {
            matrix[row][col] = array[col]
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