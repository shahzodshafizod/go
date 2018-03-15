package main

import "fmt"

func main() {
    var n int
    fmt.Print("N = ")
    fmt.Scan(&n)
    a := make([]float32, n)
    for index, _ := range a {
        fmt.Scan(&a[index])
    }
    from, to := 0, len(a) - 1
    for from < to {
        fmt.Printf("%.2f\t%.2f\t", a[from], a[to])
        from++
        to--
    }
    if from == to {
        fmt.Printf("%.2f\t", a[from])
    }
}