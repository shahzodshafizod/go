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
    var n int
    for !st.IsEmpty() {
        fmt.Printf("%d\t", st.Pop())
        n++
    }
    fmt.Printf("\nN = %d\n", n)
}