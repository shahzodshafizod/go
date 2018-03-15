package main

import "fmt"

func main() {
    var d, m int
    fmt.Print("D = ")
    fmt.Scan(&d)
    fmt.Print("M = ")
    fmt.Scan(&m)
    var zodiacalName string = ""
    switch m {
        case 1:
            switch {
                case d >= 20: zodiacalName = "Aquarius"
                default: zodiacalName = "Capricorn"
            }
        case 2:
            switch {
                case d >= 19: zodiacalName = "Pisces"
                default: zodiacalName = "Aquarius"
            }
        case 3:
            switch {
                case d >= 21: zodiacalName = "Aries"
                default: zodiacalName = "Pisces"
            }
        case 4:
            switch {
                case d >= 20: zodiacalName = "Taurus"
                default: zodiacalName = "Aries"
            }
        case 5:
            switch {
                case d >= 21: zodiacalName = "Gemini"
                default: zodiacalName = "Taurus"
            }
        case 6:
            switch {
                case d >= 22: zodiacalName = "Cancer"
                default: zodiacalName = "Gemini"
            }
        case 7:
            switch {
                case d >= 23: zodiacalName = "Leo"
                default: zodiacalName = "Cancer"
            }
        case 8:
            switch {
                case d >= 23: zodiacalName = "Virgo"
                default: zodiacalName = "Leo"
            }
        case 9:
            switch {
                case d >= 23: zodiacalName = "Libra"
                default: zodiacalName = "Virgo"
            }
        case 10:
            switch {
                case d >= 23: zodiacalName = "Scorpio"
                default: zodiacalName = "Libra"
            }
        case 11:
            switch {
                case d >= 23: zodiacalName = "Sagittarius"
                default: zodiacalName = "Scorpio"
            }
        case 12:
            switch {
                case d >= 22: zodiacalName = "Capricorn"
                default: zodiacalName = "Sagittarius"
            }
    }
    fmt.Println(zodiacalName)
}