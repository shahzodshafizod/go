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
    var index, limit int = 0, m / 2 + m % 2
    for col := 0; col < limit; col++ {
        index = m - 1 - col
        for row := col; row <= index; row++ {
            fmt.Printf("%.2f\t", a[row][col])
        }
        fmt.Println()
        for row := col+1; row <= index; row++ {
            fmt.Printf("%.2f\t", a[index][row])
        }
        fmt.Println()
        for row := index-1; row >= col; row-- {
            fmt.Printf("%.2f\t", a[row][index])
        }
        fmt.Println()
        for row := index-1; row > col; row-- {
            fmt.Printf("%.2f\t", a[col][row])
        }
        fmt.Println()
    }
}