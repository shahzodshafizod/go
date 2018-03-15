package main

import "fmt"

func main() {
    var a, b, c float32
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    fmt.Print("C = ")
    fmt.Scan(&c)
    if a < b && b < c {
        a *= 2
        b *= 2
        c *= 2
    } else {
        a, b, c = -a, -b, -c
    }
    fmt.Printf("A = %.2f\nB = %.2f\nC = %.2f\n", a, b, c)
}