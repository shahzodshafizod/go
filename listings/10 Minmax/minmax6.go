package main

import "fmt"

func main() {
    var n, number, min, max, minIndex, maxIndex int
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        fmt.Scan(&number)
        if i == 1 {
            min, max = number, number
            minIndex, maxIndex = i, i
        } else if number >= max {
            max, maxIndex = number, i
        } else if number < min {
            min, minIndex = number, i
        }
    }
    fmt.Printf("minIndex = %d\t\tmaxIndex = %d\n", minIndex, maxIndex)
}