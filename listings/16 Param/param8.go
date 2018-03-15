package main

import "fmt"

func RemoveX(a *[]int, n *int, x int) {
    for index := 0; index < *n; {
        if (*a)[index] == x {
            *a = append((*a)[:index], (*a)[index+1:]...)
            (*n)--
        } else { index++ }
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
        RemoveX(&array, &n, x)
        fmt.Println(array)
    }
}