package main

import "fmt"

func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)
    var result bool = a == b || b == c || c == a
    fmt.Printf("result = %t\n", result)
}