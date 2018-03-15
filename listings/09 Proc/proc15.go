package main

import "fmt"

/*
func ShiftLeft3(a *float32, b *float32, c *float32) {
    tmp := *c
    *c = *a
    *a = *b
    *b = tmp
}
*/

func ShiftLeft3(a float32, b float32, c float32) (float32, float32, float32) {
    return b, c, a
}

func main() {
    var a, b, c float32
    for i := 1; i <= 2; i++ {
        fmt.Printf("A%d = ", i)
        fmt.Scan(&a)
        fmt.Printf("B%d = ", i)
        fmt.Scan(&b)
        fmt.Printf("C%d = ", i)
        fmt.Scan(&c)
        //ShiftLeft3(&a, &b, &c)
        a, b, c = ShiftLeft3(a, b, c)
        fmt.Printf("\nA%d = %.2f\nB%d = %.2f\nC%d = %.2f\n\n", i, a, i, b, i, c)
    }
}