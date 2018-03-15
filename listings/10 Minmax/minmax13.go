package main

import "fmt"

func main() {
    var n, number, maxOdd, maxIndex int
    var inited bool
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        fmt.Scan(&number)
        if number % 2 != 0 {
            if !inited {
                maxOdd, maxIndex, inited = number, i, true
            } else if number > maxOdd {
                maxOdd, maxIndex = number, i
            }
        }
    }
    fmt.Printf("maxOddIndex = %d\n", maxIndex)
}