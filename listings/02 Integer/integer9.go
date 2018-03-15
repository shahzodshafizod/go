package main

import "fmt"

func main() {
    var number int
    fmt.Print("number [100-999] = ")
    fmt.Scanf("%d", &number)
    sadi := number / 100
    fmt.Printf("sadi = %d\n", sadi)
}