package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scan(&n)
    var c = string(n)
    fmt.Printf("C = %s", c)
}