package main

import "fmt"

func main() {
    var number int
    fmt.Print("number [1000:] = ")
    fmt.Scanf("%d", &number)
    sadi := number % 1000 / 100
    fmt.Printf("sadi = %d\n", sadi)
}