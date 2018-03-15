package main

import "fmt"

func main() {
    var (
        n int
        c1, c2, str string
    )
    fmt.Print("N = ")
    fmt.Scan(&n)
    fmt.Print("C1 = ")
    fmt.Scan(&c1)
    fmt.Print("C2 = ")
    fmt.Scan(&c2)
    for i := 0; i < n; i += 2 {
        str += c1 + c2
    }
    fmt.Println(str)
}