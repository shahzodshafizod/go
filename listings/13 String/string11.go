package main

import "fmt"

func main() {
    var s string
    fmt.Print("S = ")
    fmt.Scan(&s)
    var codes []rune = []rune(s)
    n := len(codes)
    var newCodes []rune = make([]rune, 2 * n - 1)
    var index int = 0
    var space rune = []rune(" ")[0]
    for i := 0; i < n; i++ {
        newCodes[index] = codes[i]
        index++
        if i < n-1 {
            newCodes[index] = space
            index++
        }
    }
    s = string(newCodes)
    fmt.Println(s)
}