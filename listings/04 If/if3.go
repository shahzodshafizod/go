package main

import "fmt"

func main() {
    var number int
    fmt.Print("number = ")
    fmt.Scanf("%d", &number)
    if number > 0 {
        number -= 8
    } else if number < 0 {
        number += 6
    } else {
        number = 10
    }
    fmt.Printf("newNumber = %d\n", number)
}