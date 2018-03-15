package main

import (
    "fmt"
    "./queue"
)

func main() {
    var (
        number int
        qu queue.TQueue
    )
    for i := 0; i < 10; i++ {
        fmt.Scan(&number)
        qu.Enqueue(number)
    }
    fmt.Printf("Head = %p\nTail = %p\n", qu.Head, qu.Tail)
    qu.Display()
}