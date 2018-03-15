package main

import (
    "fmt"
    "./stack"
)

func main() {
    var d int
    fmt.Print("D = ")
    fmt.Scan(&d)
    var st stack.TStack
    st.Make()
    st.Display()
    st.Push(d)
    st.Display()
}