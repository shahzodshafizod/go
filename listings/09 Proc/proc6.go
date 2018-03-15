package main

import "fmt"

/*
func DigitCountSum(k uint, c *uint, s *uint) {
    *c, *s = 0, 0
    for k > 0 {
        (*c)++
        *s += k % 10
        k /= 10
    }
}
*/

/*
func DigitCountSum(k uint) (c uint, s uint) {
    c, s = 0, 0
    for k > 0 {
        c++
        s += k % 10
        k /= 10
    }
    return
}
*/

func DigitCountSum(k uint) (uint, uint) {
    var c, s uint = 0, 0
    for k > 0 {
        c++
        s += k % 10
        k /= 10
    }
    return c, s
}

func main() {
    var k, count, sum uint
    for i := 0; i < 5; i++ {
        fmt.Printf("K%d = ", i+1)
        fmt.Scan(&k)
        //DigitCountSum(k, &count, &sum)
        count, sum = DigitCountSum(k)
        fmt.Printf("count = %d\t\tsum = %d\n\n", count, sum)
    }
}