package main

import "fmt"

func main() {
    var number uint
    fmt.Print("number = ")
    fmt.Scanf("%d", &number)
    var result bool = (number % 2 == 0) && (number >= 10) && (number <= 99)
    fmt.Printf("result = %t\n", result)
}