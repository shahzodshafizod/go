package main

import "fmt"

func main() {
    var a, b, c, d, index int
    fmt.Scan(&a, &b, &c, &d)
    if a == b && b == c {
        index = 4
    } else if a == b && b == d {
        index = 3
    } else if a == c && c == d {
        index = 2
    } else {
        index = 1
    }
    fmt.Printf("index = %d\n", index)
}