package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    count, sum := 0, 0
    for n > 0 {
        sum += n % 10
        n /= 10
        count++
    }
    fmt.Printf("count = %d\t\tsum = %d\n", count, sum)
}