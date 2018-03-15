package main

import "fmt"

func main() {
    var number int
    fmt.Print("number [10-99] = ")
    fmt.Scanf("%d", &number)
    dahi := number / 10
    vohid := number % 10
    sum := dahi + vohid
    mul := dahi * vohid
    fmt.Printf("sum = %d\nmultiplication = %d\n", sum, mul)
}