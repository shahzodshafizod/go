package main

import "fmt"

func main() {
    var str string
    fmt.Print("string:\t")
    fmt.Scan(&str)
    var codes []rune = []rune(str)
    fmt.Println(codes[0], codes[len(codes) - 1])
}