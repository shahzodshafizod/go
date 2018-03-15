package main

import "fmt"

func main() {
    var year int
    fmt.Print("year [20-69] = ")
    fmt.Scanf("%d", &year)
    dahi := year / 10
    vohid := year % 10
    var result string = ""
    switch dahi {
        case 2: result += "twenty"
        case 3: result += "thirty"
        case 4: result += "forty"
        case 5: result += "fifty"
        case 6: result += "sixty"
    }
    if vohid != 0 {
        result += "-"
    }
    switch vohid {
        case 1: result += "one"
        case 2: result += "two"
        case 3: result += "three"
        case 4: result += "four"
        case 5: result += "five"
        case 6: result += "six"
        case 7: result += "seven"
        case 8: result += "eight"
        case 9: result += "nine"
    }
    result += " years"
    fmt.Println(result)
}