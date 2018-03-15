package main

import "fmt"

func main() {
    var a float64
    fmt.Print("A = ")
    fmt.Scanf("%f", &a)
    var sum float64 = 0
    k := 0
    for sum < a {
        k++
        sum += 1 / float64(k)
    }
    sum -= 1 / float64(k)
    k--
    fmt.Printf("K = %d\t\tsum = %.5f\n", k, sum)
}