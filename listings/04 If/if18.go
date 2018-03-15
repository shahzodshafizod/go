package main

import "fmt"

func main() {
    var a, b, c, index int
    fmt.Scan(&a, &b, &c)
    if (a == b) {
        index = 3
    } else if a == c {
        index = 2
    } else {
        index = 1
    }
    fmt.Printf("index = %d\n", index)
}