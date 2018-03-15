package main

import "fmt"

func main() {
    var n, number, min, max, minCount, maxCount int
    fmt.Print("N = ")
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&number)
        if i == 0 {
            min, max, minCount, maxCount = number, number, 1, 1
        
        } else if number < min {
            min, minCount = number, 1
        } else if number == min {
            minCount++
        
        } else if number > max {
            max, maxCount = number, 1
        } else if number == max {
            maxCount++
        }
    }
    fmt.Printf("count = %d\n", minCount + maxCount)
}