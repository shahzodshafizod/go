package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    chappa, tmp := 0, n
    for tmp > 0 {
        chappa = chappa * 10 + tmp % 10
        tmp /= 10
    }
    fmt.Printf("chappa(%d) = %d\n", n, chappa)
}