package main

import "fmt"

func main() {
    var a [10]int
    for index, _ := range a {
        fmt.Scan(&a[index])
    }
    var nomer int
    for index := 1; index < 9; index++ {
        if a[0] < a[index] && a[index] < a[9] {
            nomer = index + 1
        }
    }
    fmt.Printf("%d\n", nomer)
}