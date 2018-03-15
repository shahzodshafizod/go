package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    sum := 0
    for i := a; i <= b; i++ {
        sum += i
    }
    fmt.Printf("sum = %d\n", sum)
}