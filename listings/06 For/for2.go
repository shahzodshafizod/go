package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    n := 0
    for i := a; i <= b; i++ {
        fmt.Printf("%d\t", i)
        n++;
    }
    fmt.Printf("\nN = %d\n", n)
}