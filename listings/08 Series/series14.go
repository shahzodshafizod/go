package main

import "fmt"

func main() {
    var k, number, count int
    fmt.Print("K = ")
    fmt.Scan(&k)
    for {
        fmt.Scan(&number)
        if number == 0 {
            break
        }
        if number < k {
            count++
        }
    }
    fmt.Printf("count = %d\n", count)
}