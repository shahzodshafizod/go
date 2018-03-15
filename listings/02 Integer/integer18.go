package main

import "fmt"

func main() {
    var number int
    fmt.Print("number [1000:] = ")
    fmt.Scanf("%d", &number)
    hazori := number / 1000 % 10
    fmt.Printf("hazori = %d\n", hazori)
}