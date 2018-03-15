package main

import "fmt"

func main() {
    var masofa float64 = 10
    var p float64
    fmt.Print("P = ")
    fmt.Scanf("%f", &p)
    days := 1
    var sum float64 = masofa
    for sum <= 200 {
        days++
        masofa += masofa / 100 * p
        sum += masofa
    }
    fmt.Printf("K = %d\t\tS = %.3f\n", days, sum)
}