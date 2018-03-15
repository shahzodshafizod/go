package main

import "fmt"

func main() {
    var k, n, number int
    fmt.Print("K = ")
    fmt.Scan(&k)
    fmt.Print("N = ")
    fmt.Scan(&n)
    lessK := false
    for n > 0 {
        fmt.Scan(&number)
        if !lessK && number < k {
            lessK = true
        }
        n--
    }
    fmt.Printf("has less then %d:\t%t\n", k, lessK)
}