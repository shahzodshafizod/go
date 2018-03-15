package main

import "fmt"

func DoubleX(a *[]int, n *int, x int) {
    for index := 0; index < *n; index++ {
        if (*a)[index] == x {
            *a = append((*a)[:index+1], (append([]int{x}, (*a)[index+1:]...))...)
            index++
            (*n)++
        }
    }
}

func main() {
    var x, n int
    for i := 1; i <= 3; i++ {
        fmt.Printf("X%d = ", i)
        fmt.Scan(&x)
        fmt.Printf("N%d = ", i)
        fmt.Scan(&n)
        var array []int = make([]int, n)
        for index, _ := range array {
            fmt.Scan(&array[index])
        }
        DoubleX(&array, &n, x)
        fmt.Println(array)
    }
}