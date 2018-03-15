package main

import "fmt"

func main() {
    var (
        b, number, min float32
        minIndex int
        inited bool
    )
    fmt.Print("B = ")
    fmt.Scan(&b)
    for i := 1; i <= 10; i++ {
        fmt.Scan(&number)
        if number > b {
            if !inited {
                min, minIndex, inited = number, i, true
            } else if number < min {
                min, minIndex = number, i
            }
        }
    }
    fmt.Printf("min = %.2f\t\tminIndex = %d\n", min, minIndex)
}