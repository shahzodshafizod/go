package main

import "fmt"

func main() {
    var number int
    fmt.Print("number [100-999] = ")
    fmt.Scanf("%d", &number)
    sadi := number / 100
    dahi := number / 10 % 10
    vohid := number % 10
    var result string = ""
    switch sadi {
        case 1: result += "one hundred "
        case 2: result += "two hundred "
        case 3: result += "three hundred "
        case 4: result += "four hundred "
        case 5: result += "five hundred "
        case 6: result += "six hundred "
        case 7: result += "seven hundred "
        case 8: result += "eight hundred "
        case 9: result += "nine hundred "
    }
    if dahi != 0 || vohid != 0 {
        result += "and "
    }
    switch dahi {
        case 1:
            switch vohid {
                case 0: result += "ten"
                case 1: result += "eleven"
                case 2: result += "twelve"
                case 3: result += "thirteen"
                case 4: result += "fourteen"
                case 5: result += "fifteen"
                case 6: result += "sixteen"
                case 7: result += "seventeen"
                case 8: result += "eighteen"
                case 9: result += "nineteen"
            }
        case 2: result += "twenty"
        case 3: result += "thirty"
        case 4: result += "forty"
        case 5: result += "fifty"
        case 6: result += "sixty"
        case 7: result += "seventy"
        case 8: result += "eighty"
        case 9: result += "ninety"
    }
    if dahi != 1 {
        if dahi > 1 {
            if vohid != 0 {
                result += "-"
            } else {
                result += " "
            }
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
    }
    fmt.Println(result)
}