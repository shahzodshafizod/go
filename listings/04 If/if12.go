package main

import "fmt"

func main() {
    var a, b, c, xurd float32
    fmt.Scan(&a, &b, &c)
    if a < b && a < c {
        xurd = a
    } else if b < c {
        xurd = b
    } else {
        xurd = c
    }
    fmt.Printf("smaller = %.2f\n", xurd)
}