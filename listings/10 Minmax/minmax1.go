package main

import "fmt"

func main() {
    var (
        n int
        number, min, max float32
        inited bool
    )
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&number)
        if !inited {
            min, max, inited = number, number, true
        } else if number < min {
            min = number
        } else if number > max {
            max = number
        }
    }
    fmt.Printf("Min = %.2f\nMax = %.2f\n", min, max)
}