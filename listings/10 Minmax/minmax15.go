package main

import "fmt"

func main() {
    var (
        b, c, number, max float32
        maxIndex int
        inited bool
    )
    fmt.Print("B = ")
    fmt.Scan(&b)
    fmt.Print("C = ")
    fmt.Scan(&c)
    for i := 1; i <= 10; i++ {
        fmt.Scan(&number)
        if b < number && number < c {
            if !inited {
                max, maxIndex, inited = number, i, true
            } else if number > max {
                max, maxIndex = number, i
            }
        }
    }
    fmt.Printf("max = %.2f\t\tminIndex = %d\n", max, maxIndex)
}