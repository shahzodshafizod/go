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
    var row int
    for col := 0; col < m; col++ {
        row = m - 1 - col;
        for i := 0; i <= row; i++ {
            fmt.Printf("%.2f\t", a[i][col])
        }
        fmt.Println()
        for i := col + 1; i < m; i++ {
            fmt.Printf("%.2f\t", a[row][i])
        }
        fmt.Println()
    }
}