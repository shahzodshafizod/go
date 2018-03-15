package main

import (
    "fmt"
    "./stack"
)

func main() {
    var st1, st2 stack.TStack
    fmt.Println("Stack #1")
    st1.Make()
    st1.Display()
    fmt.Println("Stack #2")
    st2.Make()
    st2.Display()
    fmt.Println()
    for !st1.IsEmpty() && (st1.Peek() % 2 != 0) {
        st2.Push(st1.Pop())
    }
    fmt.Printf("\nP3 = %p\n", st1.Top)
    st1.Display()
    fmt.Printf("\nP4 = %p\n", st2.Top)
    st2.Display()
}