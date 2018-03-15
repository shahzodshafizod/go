package main

import "fmt"

func main() {
    var oneKgPrice, price float64
    fmt.Print("price of one kg:\t")
    fmt.Scanf("%f", &oneKgPrice)
    for i := 0.1; i <= 1; i += 0.1 {
        price = i * oneKgPrice
        fmt.Printf("%.2f\t", price)
    }
}