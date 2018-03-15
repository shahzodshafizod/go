package main

import "fmt"

func main() {
    var n, number, min, max, minIndex, maxIndex int
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        fmt.Scan(&number)
        if i == 1 {
            min, max, minIndex, maxIndex = number, number, i, i
        } else if number >= max {
            max, maxIndex = number, i
        } else if number <= min {
            min, minIndex = number, i
        }
    }
    if minIndex > maxIndex {
        fmt.Println(minIndex)
    } else {
        fmt.Println(maxIndex)
    }
}