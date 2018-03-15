package main

import "fmt"

func main() {
    var monthNo int
    fmt.Print("monthNumber [1-12] = ")
    fmt.Scanf("%d", &monthNo)
    days := 0
    switch monthNo {
        case 2: days = 28
        case 4, 6, 9, 11: days = 30
        case 1, 3, 5, 7, 8, 10, 12: days = 31
    }
    fmt.Println("days: ", days)
}