package main

import "fmt"

func main() {
    var a, b float32;
    fmt.Print("a = ")
    fmt.Scan(&a)
    fmt.Print("b = ")
    fmt.Scan(&b)
    s := a*b
    p := 2 * (a + b)
    fmt.Printf("S = %.2f\nP = %.2f\n", s, p)
}