package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scan(&n)
    var number float32
    var sum, mul float32 = 0, 1
    for i := 0; i < n; i++ {
        fmt.Scan(&number)
        sum += number
        mul *= number
    }
    fmt.Printf("sum = %.2f\n", sum)
    fmt.Printf("multiplication = %.2f\n", mul)
}