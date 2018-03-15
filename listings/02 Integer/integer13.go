package main

import "fmt"

func main() {
    var number int
    fmt.Print("number [100-999] = ")
    fmt.Scanf("%d", &number)
    sadi := number / 100
    dahi := number / 10 % 10
    vohid := number % 10
    number = dahi * 100 + vohid * 10 + sadi
    fmt.Printf("number = %d\n", number)
}