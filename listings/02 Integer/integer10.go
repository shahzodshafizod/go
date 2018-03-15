package main

import "fmt"

func main() {
    var number int
    fmt.Print("number [100-999] = ")
    fmt.Scanf("%d", &number)
    vohid := number % 10
    dahi := number / 10 % 10
    fmt.Printf("vohid = %d\ndahi = %d\n", vohid, dahi)
}