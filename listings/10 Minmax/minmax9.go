package main

import "fmt"

func main() {
    var n, number, max, firstMaxIndex, lastMaxIndex int
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        fmt.Scan(&number)
        if i == 1 {
            max, firstMaxIndex, lastMaxIndex = number, i, i
        } else if number >= max {
            if number > max {
                max, firstMaxIndex, lastMaxIndex = number, i, i
            } else if number == max {
                lastMaxIndex = i
            }
        }
    }
    fmt.Printf("firstMaxIndex = %d\t\tlastMinIndex = %d\n", firstMaxIndex, lastMaxIndex)
}