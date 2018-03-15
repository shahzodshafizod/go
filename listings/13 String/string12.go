package main

import "fmt"

func main() {
    var s string
    var n int
    fmt.Print("S = ")
    fmt.Scan(&s)
    fmt.Print("N = ")
    fmt.Scan(&n)
    codes := []rune(s)
    darozi := len(codes)
    var newCodes []rune = make([]rune, darozi + (darozi-1)*n)
    var index int = 0
    var space rune = []rune("*")[0]
    for i := 0; i < darozi; i++ {
        newCodes[index] = codes[i]
        index++
        if i < darozi-1 {
            for j := 0; j < n; j++ {
                newCodes[index] = space
                index++
            }
        }
    }
    s = string(newCodes)
    fmt.Println(s)
}