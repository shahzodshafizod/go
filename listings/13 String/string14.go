package main

import "fmt"

func main() {
    var str string
    fmt.Print("string:\t")
    fmt.Scan(&str)
    LATS := 0
    var codes []rune = []rune(str)
    for index, _ := range codes {
        if codes[index] >= 65 && codes[index] <= 90 {
            LATS++
        }
    }
    fmt.Printf("LATINS = %d\n", LATS)
}