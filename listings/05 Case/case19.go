package main

import "fmt"

func main() {
    var year int
    const BEGINPOINT = 1984
    fmt.Scanf("%d", &year)
    var result string = "The "
    var farq, color, animal int
    if year >= BEGINPOINT {
        farq = year - BEGINPOINT
        color = farq % 60
        switch {
            case color < 12: result += "Green "
            case color < 24: result += "Red "
            case color < 36: result += "Yellow "
            case color < 48: result += "White "
            default: result += "Black "
        }
        animal = farq % 12
    } else {
        farq = BEGINPOINT - year - 1
        color = farq % 60
        switch {
            case color < 12: result += "Black "
            case color < 24: result += "White "
            case color < 36: result += "Yellow "
            case color < 48: result += "Red "
            default: result += "Green "
        }
        animal = farq % 12 + 12
    }
    switch animal {
        case 0, 23: result += "Rat"
        case 1, 22: result += "Cow"
        case 2, 21: result += "Tiger"
        case 3, 20: result += "Hare"
        case 4, 19: result += "Dragon"
        case 5, 18: result += "Snake"
        case 6, 17: result += "Horse"
        case 7, 16: result += "Sheep"
        case 8, 15: result += "Monkey"
        case 9, 14: result += "Hen"
        case 10, 13: result += "Dog"
        case 11, 12: result += "Pig"
    }
    result += "'s year"
    fmt.Println(result)
}