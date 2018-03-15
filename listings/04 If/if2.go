package main

import "fmt"

func main() {
    var number int
    fmt.Print("number = ")
    fmt.Scanf("%d", &number)
    if number > 0 {
        number -= 8
    } else {
        number += 6
    }
    fmt.Printf("newNumber = %d\n", number)
}