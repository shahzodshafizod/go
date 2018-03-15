package main

import "fmt"

/*
func Swap(x *float32, y *float32) {
    tmp := *x
    *x = *y
    *y = tmp
}
*/

func Swap(x float32, y float32) (float32, float32) {
    return y, x
}

func main() {
    var a, b, c, d float32
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    fmt.Print("C = ")
    fmt.Scan(&c)
    fmt.Print("D = ")
    fmt.Scan(&d)
    //Swap(&a, &b)
    //Swap(&c, &d)
    //Swap(&b, &c)
    a, b = Swap(a, b)
    c, d = Swap(c, d)
    b, c = Swap(b, c)
    fmt.Printf("\nA = %.2f\nB = %.2f\nC = %.2f\nD = %.2f\n", a, b, c, d)
}