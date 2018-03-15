package main

import "fmt"

func main() {
    var (
        n, minIndex int
        number, min float32
        inited bool
    )
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        fmt.Scan(&number)
        if !inited {
            min, minIndex, inited = number, i, true
        } else if number < min {
            min, minIndex = number, i
        }
    }
    fmt.Printf("minIndex = %d\n", minIndex)
}