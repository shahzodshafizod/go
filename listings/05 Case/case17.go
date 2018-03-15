package main

import "fmt"

func main() {
    var number int
    fmt.Print("number [10-40] = ")
    fmt.Scanf("%d", &number)
    dahi := number / 10
    vohid := number % 10
    var result string = "the "
    switch dahi {
        case 1:
            switch vohid {
                case 0: result += "tenth"
                case 1: result += "eleventh"
                case 2: result += "twelfth"
                case 3: result += "thirteenth"
                case 4: result += "fourteenth"
                case 5: result += "fifteenth"
                case 6: result += "sixteenth"
                case 7: result += "seventeenth"
                case 8: result += "eighteenth"
                case 9: result += "nineteenth"
            }
        case 2:
            if vohid == 0 {
                result += "twentieth"
            } else {
                result += "twenty-"
            }
        case 3:
            if vohid == 0 {
                result += "thirtieth"
            } else {
                result += "thirty-"
            }
        case 4:
            if vohid == 0 {
                result += "fortieth"
            } else {
                result += "forty-"
            }
    }
    if dahi != 1 {
        switch vohid {
            case 1: result += "first"
            case 2: result += "second"
            case 3: result += "third"
            case 4: result += "fourth"
            case 5: result += "fifth"
            case 6: result += "sixth"
            case 7: result += "seventh"
            case 8: result += "eighth"
            case 9: result += "ninth"
        }
    }
    result += " task"
    fmt.Println(result)
}