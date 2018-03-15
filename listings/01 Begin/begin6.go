package main

import "fmt"

func main() {
    var a, b, c float32
    fmt.Print("a = ")
    fmt.Scan(&a)
    fmt.Print("b = ")
    fmt.Scan(&b)
    fmt.Print("c = ")
    fmt.Scan(&c)
    v := a * b * c
    s := 2 * (a*b + b*c + a*c)
    fmt.Printf("V = %.3f\nS = %.3f\n", v, s)
}