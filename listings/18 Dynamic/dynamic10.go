package main

import (
    "fmt"
    "./stack"
)

func main() {
    var st, st1, st2 stack.TStack
    st.Make()
    st.Display()
    fmt.Println()
    for !st.IsEmpty() {
        if st.Peek() % 2 == 0 {
            st1.Push(st.Pop())
        } else {
            st2.Push(st.Pop())
        }
    }
    fmt.Printf("P2 = %p\n", st1.Top)
    st1.Display()
    fmt.Printf("P3 = %p\n", st2.Top)
    st2.Display()
}