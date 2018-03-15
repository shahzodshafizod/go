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
    for i := 0; !st.IsEmpty() && i < 5; i++ {
        fmt.Printf("%d\t", st.Pop())
    }
    fmt.Printf("\nStackIsEmpty: %t\n", st.IsEmpty())
    if !st.IsEmpty() {
        fmt.Printf("Peek: %d\n", st.Peek())
        fmt.Printf("P2 = %p\n", st.Top)
        st.Display()
    }
}