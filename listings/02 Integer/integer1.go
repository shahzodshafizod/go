package main

import "fmt"

func main() {
    var l int
    fmt.Print("L = ")
    fmt.Scanf("%d", &l)
    meters := l / 100
    fmt.Printf("meters = %d\n", meters)
}