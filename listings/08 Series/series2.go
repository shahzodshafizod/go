package main

import "fmt"

func main() {
    var number float32
    var mul float32 = 1
    for i := 0; i < 10; i++ {
        fmt.Scan(&number)
        mul *= number
    }
    fmt.Printf("multiplication = %.2f\n", mul)
}