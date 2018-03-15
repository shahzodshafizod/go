package main

import (
    "fmt"
    "time"
    "math/rand"
    "./tree"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    var tr tree.Tree
    tr.AutoMake(rand.Intn(15)+1)
    tr.Display()
    var n, index int
    fmt.Print("N = ")
    fmt.Scan(&n)
    fmt.Println()
    tr.InfixToN(&index, n)
}