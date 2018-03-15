package main

import "fmt"

/*
func PowerA234(a float32, b *float32, c *float32, d *float32) {
    *b = a * a
    *c = *b * a
    *d = *c * a
}
*/

/*
func PowerA234(a float32) (b float32, c float32, d float32) {
    b = a * a
    c = b * a
    d = c * a
    return
}
*/

func PowerA234(a float32) (float32, float32, float32) {
    b := a * a
    c := b * a
    d := c * a
    return b, c, d
}

func main() {
    var a, b, c, d float32
    for i := 0; i < 5; i++ {
        fmt.Print("A = ")
        fmt.Scan(&a)
        //PowerA234(a, &b, &c, &d)
        b, c, d = PowerA234(a)
        fmt.Printf("B = %.2f\tC = %.2f\tD = %.2f\n\n", b, c, d)
    }
}