package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    if a != b {
        kalon := a
        if a < b { kalon = b }
        a, b = kalon, kalon
    } else {
        a, b = 0, 0
    }
    fmt.Printf("A = %d\nB = %d\n", a, b)
}