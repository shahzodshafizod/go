package main

import "fmt"

func Sign(x float32) int {
    if x < 0 {
        return -1
    } else if x > 0 {
        return 1
    }
    return 0
}

func main() {
    var a, b float32
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    result := Sign(a) + Sign(b)
    fmt.Printf("result = %d\n", result)
}