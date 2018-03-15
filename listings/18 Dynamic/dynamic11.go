package main

import (
    "fmt"
    "./stack"
)

func main() {
    var st stack.TStack
    st.Make()
    st.Display()
    fmt.Print()
    var n, data int
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&data)
        st.Push(data)
    }
    fmt.Printf("\nP2 = %p\n", st.Top)
    st.Display()
}