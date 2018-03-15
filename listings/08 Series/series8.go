package main

import "fmt"

func main() {
    var n, number, k int
    fmt.Print("N = ")
    fmt.Scan(&n)
    for n > 0 {
        fmt.Scan(&number)
        if number % 2 == 0 {
            fmt.Printf("%d\t", number)
            k++
        }
        n--
    }
    fmt.Printf("\nK = %d\n", k)
}