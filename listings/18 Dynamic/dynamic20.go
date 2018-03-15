package main

import (
    "fmt"
    "./queue"
)

func main() {
    var qu queue.TQueue
    qu.Make()
    qu.Display()
    fmt.Println()
    for !qu.IsEmpty() && (qu.First() % 2 != 0) {
        fmt.Printf("%d\t", qu.Dequeue())
    }
    fmt.Printf("\nHead = %p\t\tTail = %p\n", qu.Head, qu.Tail)
    qu.Display()
}