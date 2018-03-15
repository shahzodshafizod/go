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
    for !st1.IsEmpty() {
        st2.Push(st1.Pop())
    }
    fmt.Printf("\nP3 = %p\n", st2.Top)
    st2.Display()
}