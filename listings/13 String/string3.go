package main

import "fmt"

func main() {
    var c string
    fmt.Print("C = ")
    fmt.Scan(&c)
    var code rune = []rune(c)[0]
    prev, next := string(code - 1), string(code + 1)
    fmt.Printf("Prev = %s\t\tNext = %s\n", prev, next)
}