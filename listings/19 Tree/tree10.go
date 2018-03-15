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
    var array []int = make([]int, tr.GetLevel()+1)
    tr.LevelToArray(array, 0, true)
    fmt.Println(array)
}