package main

import "fmt"

func main() {
    var a, b float32
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    if a > b {
        tmp := a
        a = b
        b = tmp
    }
    fmt.Printf("A = %.2f\nB = %.2f\n", a, b)
}