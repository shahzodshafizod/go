package main

import "fmt"

func main() {
    var k, number, index int
    fmt.Print("K = ")
    fmt.Scan(&k)
    i := 1
    for {
        fmt.Scan(&number)
        if number == 0 {
            break
        }
        if index == 0 && number > k {
            index = i
        }
        i++
    }
    fmt.Printf("index = %d\n", index)
}