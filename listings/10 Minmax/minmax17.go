package main

import "fmt"

func main() {
    var n, number, max, maxIndex int
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        fmt.Scan(&number)
        if i == 1 {
            max, maxIndex = number, i
        } else if number >= max {
            max, maxIndex = number, i
        }
    }
    fmt.Printf("elements = %d\n", n - maxIndex)
}