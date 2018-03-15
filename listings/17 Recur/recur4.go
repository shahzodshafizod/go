package main

import "fmt"

var count int

func Fib1(n int) int {
    count++
    if n <= 2 {
        return 1
    }
    return Fib1(n-1) + Fib1(n-2)
}

func main() {
    var n, element int
    for i := 1; i <= 5; i++ {
        fmt.Printf("N%d = ", i)
        fmt.Scan(&n)
        count = 0
        element = Fib1(n)
        fmt.Printf("Element = %d\t\tCount = %d\n", element, count)
    }
}