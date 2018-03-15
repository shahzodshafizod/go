package main

import "fmt"

func main() {
    var number, count int
    for {
        fmt.Scan(&number)
        if (number == 0) {
            break
        }
        count++
    }
    fmt.Printf("count = %d\n", count)
}