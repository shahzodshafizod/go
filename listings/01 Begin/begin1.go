package main

import "fmt"

func main() {
    var a float32
    fmt.Print("a = ")
    fmt.Scanf("%f", &a)
    p := a * 4
    fmt.Printf("P = %.2f", p)
}