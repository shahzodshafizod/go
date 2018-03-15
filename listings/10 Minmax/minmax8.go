package main

import "fmt"

func main() {
    var n, number, min, firstMinIndex, lastMinIndex int
    var inited bool
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        fmt.Scan(&number)
        if !inited {
            min, firstMinIndex, lastMinIndex, inited = number, i, i, true
        } else if number <= min {
            if number < min {
                min, firstMinIndex, lastMinIndex = number, i, i
            } else if number == min {
                lastMinIndex = i
            }
        }
    }
    fmt.Printf("firstMinIndex = %d\t\tlastMinIndex = %d\n", firstMinIndex, lastMinIndex)
}