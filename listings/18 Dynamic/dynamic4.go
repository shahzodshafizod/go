package main

import (
    "fmt"
    "./stack"
)

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scan(&n)
    var array []int = make([]int, n)
    for index, _ := range array {
        fmt.Scan(&array[index])
    }
    var st stack.TStack
    st.ArrayMake(array)
    fmt.Printf("P1 = %p\n", st.Top)
    st.Display()
}