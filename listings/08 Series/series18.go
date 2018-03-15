package main

import "fmt"

func main() {
    var n, number, prev int
    fmt.Print("N = ")
    fmt.Scan(&n)
    firstTime := true
    for n > 0 {
        fmt.Scan(&number)
        if firstTime {
            fmt.Printf("%d  ", number)
            firstTime = false
        } else if (number != prev) {
            fmt.Printf("%d  ", number)
        }
        prev = number
        n--
    }
}