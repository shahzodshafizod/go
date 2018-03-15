package main

import (
    "fmt"
    "./queue"
)

func main() {
    var (
        number int
        qu1, qu2 queue.TQueue
    )
    for i := 1; i <= 10; i++ {
        fmt.Scan(&number)
        if number % 2 != 0 {
            qu1.Enqueue(number)
        } else {
            qu2.Enqueue(number)
        }
    }
    fmt.Println("\nQueue #1")
    fmt.Printf("Head = %p\t\tTail = %p\n", qu1.Head, qu1.Tail)
    qu1.Display()
    fmt.Println("\nQueue #2")
    fmt.Printf("Head = %p\t\tTail = %p\n", qu2.Head, qu2.Tail)
    qu2.Display()
}