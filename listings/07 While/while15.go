package main

import "fmt"

func main() {
    var deposit float64 = 1000
    var p float64
    fmt.Print("P = ")
    fmt.Scanf("%f", &p)
    month := 0
    for deposit <= 1100 {
        month++
        deposit += deposit / 100 * p
    }
    fmt.Printf("K = %d\t\tS = %.2f\n", month, deposit)
}