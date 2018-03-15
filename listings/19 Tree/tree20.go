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
    fmt.Println()
    minimal := tr.GetMinData()
    count := tr.GetLeafDataCount(minimal)
    fmt.Printf("minimal = %d\t\tcount = %d\n", minimal, count)
}