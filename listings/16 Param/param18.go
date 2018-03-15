package main

import "fmt"

func Chessboard(m int, n int, a *[][]int) {
    (*a) = make([][]int, m)
    for row, _ := range *a {
        (*a)[row] = make([]int, n)
        for col, _ := range (*a)[row] {
            if (row+col) % 2 == 0 {
                (*a)[row][col] = 0
            } else {
                (*a)[row][col] = 1
            }
        }
    }
}

func main() {
    var m, n int
    fmt.Print("M = ")
    fmt.Scan(&m)
    fmt.Print("N = ")
    fmt.Scan(&n)
    var matrix [][]int;
    Chessboard(m, n, &matrix)
    for row, _ := range matrix {
        fmt.Println(matrix[row])
    }
}