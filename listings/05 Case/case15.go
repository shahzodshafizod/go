package main

import "fmt"

func main() {
    var n, m int
    fmt.Print("N = ")
    fmt.Scan(&n)
    fmt.Print("M = ")
    fmt.Scan(&m)
    var result string = ""
    switch n {
        case 6: result += "six"
        case 7: result += "seven"
        case 8: result += "eight"
        case 9: result += "nine"
        case 10: result += "ten"
        case 11: result += "jack"
        case 12: result += "queen"
        case 13: result += "king"
        case 14: result += "ace"
    }
    result += " of "
    switch m {
        case 1: result += "spades"
        case 2: result += "clubs"
        case 3: result += "diamonds"
        case 4: result += "hearts"
    }
    fmt.Println(result)
}