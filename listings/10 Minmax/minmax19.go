package main

import "fmt"

func main() {
    var n, number, min, count int
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&number)
        if i == 0 {
            min, count = number, 1
        } else if number < min {
            min, count = number, 1
        } else if number == min {
            count++
        }
    }
    fmt.Printf("count = %d\n", count)
}