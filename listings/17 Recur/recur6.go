package main

import "fmt"

var count int

func Combin1(n, k int) int {
    count++
    if k == 0 || k == n {
        return 1
    }
    return Combin1(n-1, k) + Combin1(n-1, k-1)
}

func main() {
    var n, k, comb int
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 1; i <= 5; i++ {
        fmt.Printf("K%d = ", i)
        fmt.Scan(&k)
        count = 0
        comb = Combin1(n, k)
        fmt.Printf("Combin(%d, %d) = %d\t\tCount = %d\n", n, k, comb, count)
    }
}