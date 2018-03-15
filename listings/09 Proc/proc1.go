package main

import "fmt"

func PowerA3(a float32, b *float32) {
    *b = a * a * a
}

func main() {
    var a, b float32
    for i := 0; i < 5; i++ {
        fmt.Print("A = ")
        fmt.Scan(&a)
        PowerA3(a, &b)
        fmt.Printf("B = %.2f\n\n", b)
    }
}