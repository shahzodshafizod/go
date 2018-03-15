package main

import "fmt"

var array [20]int

func Fib2(n int) int {
    if n <= 2 {
        return 1
    }
    if array[n-1] == 0 {
        array[n - 1] = Fib2(n-1) + Fib2(n-2)
    }
    return array[n-1]
}

func main() {
    var n, element int
    for i := 1; i <= 5; i++ {
        fmt.Printf("N%d = ", i)
        fmt.Scan(&n)
        for index, _ := range array {
            array[index] = 0
        }
        element = Fib2(n)
        fmt.Printf("Fib2(%d) = %d\n", n, element)
    }
}