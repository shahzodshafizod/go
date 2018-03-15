package main

import "fmt"

func main() {
    var (
        n int
        number, minPositive float32
        inited bool
    )
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        fmt.Scan(&number)
        if number > 0 {
            if !inited {
                minPositive, inited = number, true
            } else if number < minPositive {
                minPositive = number
            }
        }
    }
    fmt.Printf("minPositive = %.2f\n", minPositive)
}