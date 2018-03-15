package main

import "fmt"

func main() {
    var n, number int
    fmt.Print("N = ")
    fmt.Scan(&n)
    hasPositive := false
    for n > 0 {
        fmt.Scan(&number)
        if !hasPositive && number > 0 {
            hasPositive = true
        }
        n--
    }
    fmt.Printf("has positive:\t%t\n", hasPositive)
}