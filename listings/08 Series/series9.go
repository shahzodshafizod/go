package main

import "fmt"

func main() {
    var n, number, k int
    fmt.Print("N = ")
    fmt.Scan(&n)
    index := 1
    for index <= n {
        fmt.Scan(&number)
        if number % 2 != 0 {
            fmt.Printf("%d\t", index)
            k++
        }
        index++
    }
    fmt.Printf("\nK = %d\n", k)
}