package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    mul := 1
    for i := a; i <= b; i++ {
        mul *= i
    }
    fmt.Printf("multiplication = %d\n", mul)
}