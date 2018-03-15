package main

import "fmt"

func GCD(a, b int) int {
    if b == 0 {
        return a
    }
    return GCD(b, a % b)
}

func main() {
    var a, b, c, d int
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    fmt.Print("C = ")
    fmt.Scan(&c)
    fmt.Print("D = ")
    fmt.Scan(&d)
    fmt.Printf("GCD(%d, %d) = %d\n", a, b, GCD(a, b))
    fmt.Printf("GCD(%d, %d) = %d\n", a, c, GCD(a, c))
    fmt.Printf("GCD(%d, %d) = %d\n", a, d, GCD(a, d))
}