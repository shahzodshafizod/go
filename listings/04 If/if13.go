package main

import "fmt"

func main() {
    var a, b, c, kalon, xurd float32
    fmt.Scan(&a, &b, &c)
    if a < b && a < c {
        xurd = a
    } else if b < c {
        xurd = b
    } else {
        xurd = c
    }
    if a > b && a > c {
        kalon = a
    } else if b > c {
        kalon = b
    } else {
        kalon = c
    }
    bayn := a + b + c - kalon - xurd
    fmt.Printf("bayn = %.2f", bayn)
}