package main

import "fmt"

func main() {
    var number, sum int
    for {
        fmt.Scan(&number)
        if (number == 0) {
            break
        }
        if number > 0 && (number % 2 == 0) {
            sum += number
        }
    }
    fmt.Printf("sum = %d\n", sum)
}