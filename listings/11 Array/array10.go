package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scan(&n)
    var array []int = make([]int, n)
    for index, _ := range array {
        fmt.Scan(&array[index])
    }
    for _, value := range array {
        if value % 2 == 0 {
            fmt.Printf("%d\t", value)
        }
    }
    for index := len(array) - 1; index >= 0; index-- {
        if array[index] % 2 != 0 {
            fmt.Printf("%d\t", array[index])
        }
    }
}