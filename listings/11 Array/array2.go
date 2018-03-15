package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    var arr []int = make([]int, n)
    for index, _ := range arr {
        if index == 0 {
            arr[index] = 2
        } else {
            arr[index] = arr[index-1] + arr[index-1]
        }
    }
    fmt.Println(arr)
}