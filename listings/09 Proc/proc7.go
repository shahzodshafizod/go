package main

import "fmt"

/*
func InvDigits(k *uint) {
    tmp := *k
    *k = 0
    for tmp > 0 {
        *k = (*k) * 10 + tmp % 10
        tmp /= 10
    }
}
*/

/*
func InvDigits(k uint) uint {
    var tmp uint = 0
    for k > 0 {
        tmp = tmp * 10 + k % 10
        k /= 10
    }
    return tmp
}
*/

func InvDigits(k uint) (tmp uint) {
    tmp = 0
    for k > 0 {
        tmp = tmp * 10 + k % 10
        k /= 10
    }
    return
}

func main() {
    var k uint
    for i := 1; i <= 5; i++ {
        fmt.Printf("K%d = ", i)
        fmt.Scan(&k)
        //InvDigits(&k)
        k = InvDigits(k)
        fmt.Printf("K%d = %d\n\n", i, k)
    }
}