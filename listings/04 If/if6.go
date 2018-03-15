package main

import "fmt"

func main() {
    var a, b, kalon float64
    fmt.Scan(&a, &b)
    if a > b {
        kalon = a
    } else {
        kalon = b
    }
    fmt.Printf("greater is %.2f\n", kalon)
}