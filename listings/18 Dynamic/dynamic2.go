package main

import (
    "fmt"
    "./stack"
)

func main() {
    var st *stack.TStack = new(stack.TStack)
    st.Make()
    st.Display()
    fmt.Println()
    var (
        n int
        lastNode *stack.TNode
    )
    for i := st.Top; i != nil; i = i.Next {
        n++
        fmt.Printf("%d\t", i.Data)
        lastNode = i
    }
    fmt.Printf("\nN = %d\n", n)
    fmt.Printf("lastNode = %p\n", lastNode)
    st.Display()
}