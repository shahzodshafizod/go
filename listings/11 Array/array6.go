package main

import "fmt"

func main() {
    var n, a, b int
    fmt.Print("N = ")
    fmt.Scan(&n)
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    var array []int = make([]int, n)
    array[0], array[1] = a, b
    var sum int = a + b
    for index := 2; index < n; index++ {
        array[index] = sum
        sum += array[index]
    }
    fmt.Println(array)
}