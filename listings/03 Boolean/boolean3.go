package main

import "fmt"

func main() {
    var a int
    fmt.Print("A = ")
    fmt.Scanf("%d", &a)
    isEven := a % 2 == 0
    fmt.Printf("Even = %t\n", isEven)
}