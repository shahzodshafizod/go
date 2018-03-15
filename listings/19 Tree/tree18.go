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
    var level int
    fmt.Print("L = ")
    fmt.Scan(&level)
    fmt.Println()
    n := tr.PrintLevel(0, level)
    fmt.Printf("\nN = %d\n", n)
}