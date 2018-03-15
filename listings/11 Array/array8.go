package main

import "fmt"

func main() {
    var n, k int
    fmt.Print("N = ")
    fmt.Scan(&n)
    var array []int = make([]int, n)
    for index, _ := range array {
        fmt.Scan(&array[index])
    }
    for _, value := range array {
        if value % 2 != 0 {
            fmt.Printf("%d\t", value)
            k++
        }
    }
    fmt.Printf("\nK = %d\n", k)
}