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
    for {
        if from > to { break }
        fmt.Printf("%.2f\t", a[from])
        from++
        
        if from > to { break }
        fmt.Printf("%.2f\t", a[from])
        from++
        
        if from > to { break }
        fmt.Printf("%.2f\t", a[to])
        to--
        
        if from > to { break }
        fmt.Printf("%.2f\t", a[to])
        to--
    }
}