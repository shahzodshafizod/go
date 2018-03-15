package main

import "fmt"

func main() {
    var a int
    fmt.Print("A = ")
    fmt.Scanf("%d", &a)
    isOdd := a % 2 != 0
    fmt.Printf("Odd = %t\n", isOdd)
}