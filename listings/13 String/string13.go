package main

import "fmt"

func main() {
    var str string
    fmt.Print("string:\t")
    fmt.Scan(&str)
    digits := 0
    var codes []rune = []rune(str)
    for index, _ := range codes {
        if codes[index] >= 48 && codes[index] <= 57 {
            digits++
        }
    }
    fmt.Printf("digits = %d\n", digits)
}