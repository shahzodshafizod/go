package main

import "fmt"

func main() {
    var a, b float32
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    freeSpace := a
    for freeSpace >= b {
        freeSpace -= b
    }
    fmt.Printf("free space: %.2f\n", freeSpace)
}