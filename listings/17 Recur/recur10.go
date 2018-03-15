package main

import "fmt"

func DigitSum(number int) int {
    if number == 0 {
        return 0
    }
    if number < 0 {
        number *= -1
    }
    return number % 10 + DigitSum(number / 10)
}

func main() {
    var k int
    for i := 1; i <= 5; i++ {
        fmt.Printf("K%d = ", i)
        fmt.Scan(&k)
        fmt.Printf("DigitSum(%d) = %d\n", k, DigitSum(k))
    }
}