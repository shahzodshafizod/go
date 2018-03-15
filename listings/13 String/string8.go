package main

import "fmt"

func main() {
    var (
        n int
        c, str string
    )
    fmt.Print("N = ")
    fmt.Scan(&n)
    fmt.Print("C = ")
    fmt.Scan(&c)
    for i := 0; i < n; i++ {
        str += c
    }
    fmt.Println(str)
}