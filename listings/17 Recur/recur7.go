package main

import "fmt"

var matrix [20][20]int

func Combin2(n, k int) int {
    if k == 0 || k == n {
        return 1
    }
    if (matrix[n-1][k-1] == 0) {
        matrix[n-1][k-1] = Combin2(n-1, k) + Combin2(n-1, k-1)
    }
    return matrix[n-1][k-1]
}

func main() {
    var n, k, comb int
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 1; i <= 5; i++ {
        fmt.Printf("K%d = ", i)
        fmt.Scan(&k)
        for row, _ := range matrix {
            for col, _ := range matrix[row] {
                matrix[row][col] = 0
            }
        }
        comb = Combin2(n, k)
        fmt.Printf("Combin(%d, %d) = %d\n", n, k, comb)
    }
}