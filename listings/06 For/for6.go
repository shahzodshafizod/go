package main

import "fmt"

func main() {
    var oneKgPrice, price float64
    fmt.Print("price of one kg:\t")
    fmt.Scanf("%f", &oneKgPrice)
    for i := 1.2; i <= 2; i += 0.2 {
        price = i * oneKgPrice
        fmt.Printf("%.2f\t", price)
    }
}