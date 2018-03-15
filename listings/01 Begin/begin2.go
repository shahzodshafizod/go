package main

import "fmt"

func main() {
    var a float32
    fmt.Print("a = ")
    fmt.Scanf("%f", &a)
    s := a*a
    fmt.Printf("S = %.2f", s)
}