package main

import "fmt"

func toLowerLat(str string) string {
    var codes []rune = []rune(str)
    const (
        GREAT_A rune = 65
        GREAT_Z rune = 90
        SMALL_A rune = 97
        SMALL_Z rune = 122
    )
    var farq rune
    for index, _ := range codes {
        if codes[index] >= GREAT_A && codes[index] <= GREAT_Z {
            farq = codes[index] - GREAT_A
            codes[index] = SMALL_A + farq
        }
    }
    return string(codes)
}

func main() {
    var str string
    fmt.Print("string:\t")
    fmt.Scan(&str)
    str = toLowerLat(str)
    fmt.Printf("string = %s\n", str)
}