package main

import "fmt"

func main() {
    var n, number, min, minIndex int
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        fmt.Scan(&number)
        if i == 1 {
            min, minIndex = number, i
        } else if number < min {
            min, minIndex = number, i
        }
    }
    fmt.Printf("elements = %d\n", minIndex-1)
}