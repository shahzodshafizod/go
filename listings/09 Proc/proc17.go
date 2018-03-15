package main

import "fmt"

func RootCount(a float32, b float32, c float32) int {
    d := b*b - 4*a*c
    roots := 0
    if d > 0 {
        roots += 2
    } else if d == 0 {
        roots++
    }
    return roots
}

func main() {
    var a, b, c float32
    for i := 1; i <= 3; i++ {
        fmt.Printf("A%d = ", i)
        fmt.Scan(&a)
        fmt.Printf("B%d = ", i)
        fmt.Scan(&b)
        fmt.Printf("C%d = ", i)
        fmt.Scan(&c)
        fmt.Printf("RootCount(A%d, B%d, C%d) = %d\n\n", i, i, i, RootCount(a, b, c))
    }
}