package main

import "fmt"

func main() {
    var number int
    fmt.Print("number [10-99] = ")
    fmt.Scanf("%d", &number)
    dahi := number / 10
    vohid := number % 10
    number = vohid * 10 + dahi
    fmt.Printf("number = %d\n", number)
}