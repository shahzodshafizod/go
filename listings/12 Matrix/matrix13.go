package main

import "fmt"

func main() {
    var m int
    fmt.Print("M = ")
    fmt.Scan(&m)
    var a [][]float32 = make([][]float32, m)
    for row, _ := range a {
        a[row] = make([]float32, m)
        for col, _ := range a[row] {
            fmt.Scan(&a[row][col])
        }
    }
    fmt.Println()
    var col int
    for row, _ := range a {
        col = m - 1 - row;
        for j := 0; j <= col; j++ {
            fmt.Printf("%.2f\t", a[row][j])
        }
        fmt.Println()
        for j := row + 1; j < m; j++ {
            fmt.Printf("%.2f\t", a[j][col])
        }
        fmt.Println()
    }
}