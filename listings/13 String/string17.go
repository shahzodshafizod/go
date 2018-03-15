package main

import "fmt"

func toUpper(str string) string {
    var codes []rune = []rune(str)
    const (
        GREAT_LAT_A, GREAT_LAT_Z rune = 65, 90
        SMALL_LAT_A, SMALL_LAT_Z rune = 97, 122
        
        GREAT_RUS_A, GREAT_RUS_Z rune = 1040, 1071
        SMALL_RUS_A, SMALL_RUS_Z rune = 1072, 1103
    )
    var farq rune
    for index, _ := range codes {
        if codes[index] >= SMALL_LAT_A && codes[index] <= SMALL_LAT_Z {
            farq = codes[index] - SMALL_LAT_A
            codes[index] = GREAT_LAT_A + farq
        } else if codes[index] >= SMALL_RUS_A && codes[index] <= SMALL_RUS_Z {
            farq = codes[index] - SMALL_RUS_A
            codes[index] = GREAT_RUS_A + farq
        }
    }
    return string(codes)
}

func main() {
    var str string
    fmt.Print("string:\t")
    fmt.Scan(&str)
    str = toUpper(str)
    fmt.Printf("string = %s\n", str)
}