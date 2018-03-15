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
    var k int
    fmt.Print("K = ")
    fmt.Scan(&k)
    fmt.Printf("\n%d\n", tr.GetDataCount(k))
}