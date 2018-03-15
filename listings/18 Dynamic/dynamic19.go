package main

import (
    "fmt"
    "./queue"
)

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scan(&n)
    var qu queue.TQueue
    qu.Make()
    qu.Display()
    fmt.Println()
    for i := 0; !qu.IsEmpty() && i < n; i++ {
        fmt.Printf("%d\t", qu.Dequeue())
    }
    fmt.Printf("\nHead = %p\t\tTail = %p\n", qu.Head, qu.Tail)
    qu.Display()
}