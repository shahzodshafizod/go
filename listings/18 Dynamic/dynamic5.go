package main

import (
    "fmt"
    "./stack"
)

func main() {
    var st stack.TStack
    st.Make()
    st.Display()
    fmt.Println()
    d := st.Pop()
    fmt.Printf("D = %d\n", d)
    fmt.Printf("P2 = %p\n", st.Top)
    st.Display()
}