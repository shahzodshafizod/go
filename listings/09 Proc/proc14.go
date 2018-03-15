package main

import "fmt"

/*
func ShiftRight3(a *float32, b *float32, c *float32) {
    tmp := *c
    *c = *b
    *b = *a
    *a = tmp
}
*/

func ShiftRight3(a float32, b float32, c float32) (float32, float32, float32) {
    return c, a, b
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
        //ShiftRight3(&a, &b, &c)
        a, b, c = ShiftRight3(a, b, c)
        fmt.Printf("\nA%d = %.2f\nB%d = %.2f\nC%d = %.2f\n\n", i, a, i, b, i, c)
    }
}