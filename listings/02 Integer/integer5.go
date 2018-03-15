package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    freeSpace := a % b
    fmt.Printf("freeSpace = %d\n", freeSpace)
}