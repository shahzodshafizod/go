package main

import "fmt"

/*
func Swap(x *float32, y *float32) {
    tmp := *x
    *x = *y
    *y = tmp
}
func Minmax(x *float32, y *float32) {
    if *x > *y {
        Swap(x, y)
    }
}
*/

/*
func Minmax(x float32, y float32) (float32, float32) {
    if x > y {
        return y, x
    }
    return x, y
}
*/

func Minmax(x float32, y float32) (min float32, max float32) {
    if x < y {
        min, max = x, y
    } else {
        min, max = y, x
    }
    return
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
    //Minmax(&a, &b)
    //Minmax(&c, &d)
    //Minmax(&a, &c)
    //Minmax(&b, &d)
    a, b = Minmax(a, b)
    c, d = Minmax(c, d)
    a, c = Minmax(a, c)
    b, d = Minmax(b, d)
    fmt.Printf("Min = %.2f\t\tMax = %.2f\n", a, d)
}