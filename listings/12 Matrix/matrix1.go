package main

import "fmt"

func main() {
    var m, n int
    fmt.Print("M = ")
    fmt.Scan(&m)
    fmt.Print("N = ")
    fmt.Scan(&n)
    var matrix [][]int = make([][]int, m)
    for row, _ := range matrix {
        matrix[row] = make([]int, n)
        for col, _ := range matrix[row] {
            matrix[row][col] = 10 * (row + 1)
        }
    }
    for _, array := range matrix {
        for _, value := range array {
            fmt.Printf("%d\t", value)
        }
        fmt.Println()
    }
}