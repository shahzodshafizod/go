package main

import "fmt"

func main() {
    var oneKgPrice float32
    fmt.Print("price of one kg:  ")
    fmt.Scanf("%f", &oneKgPrice)
    price := oneKgPrice
    for i := 1; i <= 10; i++ {
        fmt.Printf("%.2f\t", price)
        price += oneKgPrice
    }
}