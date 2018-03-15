package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    n := 0
    for i := b-1; i > a; i-- {
        fmt.Printf("%d\t", i)
        n++
    }
    fmt.Printf("\nN = %d\n", n)
}