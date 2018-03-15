package main

import "fmt"

func main() {
    var str string
    fmt.Print("string:\t")
    fmt.Scan(&str)
    var codes []rune = []rune(str)
    var point rune = []rune(".")[0]
    var index, darozi = 0, len(codes)
    var negative, points, digits int
    if codes[0] == []rune("-")[0] {
        index++
        negative = 1
    }
    for index < darozi {
        if codes[index] == point {
            points++
        } else if codes[index] >= 48 && codes[index] <= 57 {
            digits++
        }
        index++
    }
    if negative + digits == darozi {
        fmt.Println("1")
    } else if points == 1 && (negative + digits + points == darozi) {
        fmt.Println("2")
    } else {
        fmt.Println("0")
    }
}