package main

import "fmt"

func main() {
    var n, k int
    fmt.Print("N = ")
    fmt.Scan(&n)
    fmt.Print("K = ")
    fmt.Scan(&k)
    div := 0
    for n >= k {
        n -= k
        div++
    }
    mod := n
    fmt.Printf("div = %d\nmod = %d\n", div, mod)
}