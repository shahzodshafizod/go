package main

import "fmt"

func main() {
    var a, b float32
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    porchaho := 0
    for a >= b {
        a -= b
        porchaho++
    }
    fmt.Printf("porchaho = %d\n", porchaho)
}