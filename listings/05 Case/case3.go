package main

import "fmt"

func main() {
    var monthNo int
    fmt.Print("monthNumber [1-12] = ")
    fmt.Scanf("%d", &monthNo)
    seasonName := ""
    switch monthNo {
        case 1, 2, 12: seasonName = "Winter"
        case 3, 4, 5: seasonName = "Spring"
        case 6, 7, 8: seasonName = "Summer"
        case 9, 10, 11: seasonName = "Autumn"
    }
    fmt.Println("seasonName = ", seasonName)
}