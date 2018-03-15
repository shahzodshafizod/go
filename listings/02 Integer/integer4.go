package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    porchaho := a / b;
    fmt.Printf("porchaho = %d\n", porchaho)
}