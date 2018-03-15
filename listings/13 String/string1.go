package main

import "fmt"

func main() {
    var c string;
    fmt.Print("C = ")
    fmt.Scan(&c)
    var n rune = []rune(c)[0]
    fmt.Printf("N = %d\n", n)
}