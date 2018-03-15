package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scanf("%d", &n)
    var array []int = make([]int, n)
    array[0], array[1] = 1, 1
    for index := 2; index < n; index++ {
        array[index] = array[index-1] + array[index-2]
    }
    fmt.Println(array)
}