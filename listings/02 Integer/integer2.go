package main

import "fmt"

func main() {
    var m int
    fmt.Print("M = ")
    fmt.Scanf("%d", &m)
    tons := m / 1000
    fmt.Printf("tons = %d\n", tons)
}