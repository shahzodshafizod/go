package main

import (
    "fmt"
    "math/rand"
    "time"
    "./tree"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    var tr tree.Tree
    tr.AutoMake(rand.Intn(15)+1)
    tr.Display()
    var n1, n2, index int
    fmt.Print("N1 = ")
    fmt.Scan(&n1)
    fmt.Print("N2 = ")
    fmt.Scan(&n2)
    fmt.Println()
    tr.PrefixBetween(&index, n1, n2)
}