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
    for i := 0; i < 5; i++ {
        fmt.Printf("%d\t", st.Pop())
    }
    fmt.Printf("\nP2 = %p\n", st.Top)
    st.Display()
}