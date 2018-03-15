package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    hours := n / 3600
    fmt.Printf("hours = %d\n", hours)
}