package main

import "fmt"

/*
func AddRightDigit(d uint, k *uint) {
    *k = (*k) * 10 + d
}
*/

/*
func AddRightDigit(d uint, k uint) uint {
    return k * 10 + d
}
*/

func AddRightDigit(d uint, k uint) (tmp uint) {
    tmp = k * 10 + d
    return
}

func main() {
    var k, d uint
    fmt.Print("K = ")
    fmt.Scan(&k)
    for i := 1; i <= 2; i++ {
        fmt.Printf("D%d = ", i)
        fmt.Scan(&d)
        //AddRightDigit(d, &k)
        k = AddRightDigit(d, k)
        fmt.Printf("K = %d\n\n", k)
    }
}