package main

import "fmt"

func main() {
    var str string
    fmt.Print("string:\t")
    fmt.Scan(&str)
    letters := 0
    var codes []rune = []rune(str)
    for index, _ := range codes {
        if codes[index] >= 1072 && codes[index] <= 1103 ||
           codes[index] >= 97 && codes[index] <= 122 {
            letters++
        }
    }
    fmt.Printf("letters = %d\n", letters)
}