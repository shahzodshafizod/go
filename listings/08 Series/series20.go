package main

import "fmt"

func main() {
    var n, number, prev, k int
    fmt.Print("N = ")
    fmt.Scan(&n)
    fmt.Scan(&prev)
    for n > 1 {
        fmt.Scan(&number)
        if prev < number {
            fmt.Printf("%d  ", prev)
            k++
        }
        prev = number
        n--
    }
    fmt.Printf("\nK = %d\n", k)
}