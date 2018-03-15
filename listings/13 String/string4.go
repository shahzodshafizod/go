package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    code := 65
    for i := 0; i < n; i++ {
        fmt.Printf("%s\t", string(code + i))
    }
}