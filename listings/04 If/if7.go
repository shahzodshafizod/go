package main

import "fmt"

func main() {
    var a, b float64
    var index uint
    fmt.Scan(&a, &b)
    if a < b {
        index = 1
    } else {
        index = 2
    }
    fmt.Printf("index = %d\n", index)
}