package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    minutes := n / 60
    fmt.Printf("minutes = %d\n", minutes)
}