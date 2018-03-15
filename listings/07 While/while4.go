package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    degree3 := 1
    for degree3 < n {
        degree3 *= 3
    }
    result := degree3 == n
    fmt.Printf("result = %t\n", result)
}