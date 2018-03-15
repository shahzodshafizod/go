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
    for row := 0; row < limit; row++ {
        index = m - row - 1
        for col := row; col <= index; col++ {
            fmt.Printf("%.2f\t", a[row][col])
        }
        fmt.Println()
        for col := row+1; col <= index; col++ {
            fmt.Printf("%.2f\t", a[col][index])
        }
        fmt.Println()
        for col := index-1; col >= row; col-- {
            fmt.Printf("%.2f\t", a[index][col])
        }
        fmt.Println()
        for col := index-1; col > row; col-- {
            fmt.Printf("%.2f\t", a[col][row])
        }
        fmt.Println()
    }
}