package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    has2 := false
    for n > 0 {
        if n % 10 == 2 {
            has2 = true
            break
        }
        n /= 10
    }
    fmt.Printf("has2 = %t\n", has2)
}