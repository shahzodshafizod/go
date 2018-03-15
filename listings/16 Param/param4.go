package main

import "fmt"

func Inv(a []float32, n int) {
    from, to := 0, n - 1
    for from < to {
        tmp := a[from]
        a[from] = a[to]
        a[to] = tmp
        from++
        to--
    }
}

func main() {
    var n int
    for i := 1; i <= 3; i++ {
        fmt.Printf("N%d = ", i)
        fmt.Scan(&n)
        var array []float32 = make([]float32, n)
        for j := 0; j < n; j++ {
            fmt.Scan(&array[j])
        }
        Inv(array, n)
        for index, _ := range array {
            fmt.Printf("%.2f\t", array[index])
        }
        fmt.Println("\n")
    }
}