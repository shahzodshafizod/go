package main

import "fmt"

func main() {
    var number int
    fmt.Print("number [1-7] = ")
    fmt.Scanf("%d", &number)
    weekDay := "Errorday";
    switch number {
        case 1: weekDay = "Monday"
        case 2: weekDay = "Tuesday"
        case 3: weekDay = "Wednesday"
        case 4: weekDay = "Thursday"
        case 5: weekDay = "Friday"
        case 6: weekDay = "Saturday"
        case 7: weekDay = "Sunday"
    }
    fmt.Printf("weekDay: %s\n", weekDay)
}