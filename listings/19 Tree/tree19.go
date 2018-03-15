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
    maximal := tr.GetMaxData()
    count := tr.GetDataCount(maximal)
    fmt.Printf("maximal = %d\t\tcount = %d\n", maximal, count)
}