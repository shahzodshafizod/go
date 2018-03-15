package main

import (
    "fmt"
    "./queue"
)

func main() {
    var (
        d int
        qu queue.TQueue
    )
    fmt.Print("D = ")
    fmt.Scan(&d)
    qu.Make()
    qu.Display()
    fmt.Println()
    qu.Enqueue(d)
    fmt.Printf("Dequeued: %d\n", qu.Dequeue())
    fmt.Printf("P3 = %p\t\tP4 = %p\n", qu.Head, qu.Tail)
    qu.Display()
}