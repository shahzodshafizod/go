package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scan(&n)
    arr := make([]int, n)
    for index, _ := range arr {
        arr[index] = 2 * index + 1
    }
    fmt.Println(arr)
}