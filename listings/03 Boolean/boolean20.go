package main

import "fmt"

func main() {
    var number int
    fmt.Print("number [100-999] = ")
    fmt.Scanf("%d", &number)
    var sadi int = number / 100
    var dahi int = number / 10 % 10
    var vohid int = number % 10
    var result bool = sadi != dahi && dahi != vohid && vohid != sadi
    fmt.Printf("result = %t\n", result)
}