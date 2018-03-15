package main

import "fmt"

/*
func Swap(x *float32, y *float32) {
    tmp := *x
    *x = *y
    *y = tmp
}
func Maxmin(x *float32, y *float32) {
    if *x < *y {
        Swap(x, y)
    }
}
func SortDec3(a *float32, b *float32, c *float32) {
    Maxmin(a, b)
    Maxmin(a, c)
    Maxmin(b, c)
}
*/

func Maxmin(x float32, y float32) (float32, float32) {
    if x < y {
        return y, x
    }
    return x, y
}
/*
func SortDec3(a float32, b float32, c float32) (x float32, y float32, z float32) {
    a, b = Maxmin(a, b)
    a, c = Maxmin(a, c)
    b, c = Maxmin(b, c)
    x, y, z = a, b, c
    return
}
*/
func SortDec3(a float32, b float32, c float32) (float32, float32, float32) {
    a, b = Maxmin(a, b)
    a, c = Maxmin(a, c)
    b, c = Maxmin(b, c)
    return a, b, c
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
        //SortDec3(&a, &b, &c)
        a, b, c = SortDec3(a, b, c)
        fmt.Printf("A%d = %.2f\tB%d = %.2f\tC%d = %.2f\n\n", i, a, i, b, i, c)
    }
}