package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    for n > 0 {
        fmt.Printf("%d\t", n % 10)
        n /= 10
    }
}