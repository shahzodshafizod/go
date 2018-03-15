package main

import "fmt"

func main() {
    var a [10]int
    for index, _ := range a {
        fmt.Scan(&a[index])
    }
    var value int
    for index := 0; index < 9; index++ {
        if a[index] < a[9] {
            value = a[index]
            break
        }
    }
    fmt.Printf("%d\n", value)
}