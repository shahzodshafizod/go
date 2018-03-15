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
    var array []float32 = make([]float32, m)
    for index, _ := range array {
        fmt.Scan(&array[index])
    }
    var matrix [][]float32 = make([][]float32, m)
    for row, _ := range matrix {
        matrix[row] = make([]float32, n)
        for col, _ := range matrix[row] {
            if col == 0 {
                matrix[row][col] = array[row]
            } else {
                matrix[row][col] = matrix[row][col-1] + d
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