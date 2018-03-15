package main

import "fmt"

/*
func AddLeftDigit(d uint, k *uint) {
    tmp := *k
    for tmp > 0 {
        d *= 10
        tmp /= 10
    }
    *k += d
}
*/

/*
func AddLeftDigit(d uint, k uint) uint {
    tmp := k
    for tmp > 0 {
        d *= 10
        tmp /= 10
    }
    return k + d
}
*/

func AddLeftDigit(d uint, k uint) (tmp uint) {
    tmp = k
    for k > 0 {
        d *= 10
        k /= 10
    }
    tmp += d
    return
}

func main() {
    var k, d uint
    fmt.Print("K = ")
    fmt.Scan(&k)
    for i := 1; i <= 2; i++ {
        fmt.Printf("D%d = ", i)
        fmt.Scan(&d)
        //AddLeftDigit(d, &k)
        k = AddLeftDigit(d, k)
        fmt.Printf("K = %d\n\n", k)
    }
}