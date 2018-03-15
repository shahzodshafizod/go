package main

import "fmt"

func main() {
    var a, b, kalon, xurd float64
    fmt.Scan(&a, &b)
    kalon, xurd = a, b
    if kalon < xurd {
        kalon, xurd = b, a
    }
    fmt.Printf("greater = %.2f\nsmaller = %.2f\n", kalon, xurd)
}