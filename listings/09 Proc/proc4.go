package main

import (
    "fmt"
    "math"
)

/*
func TrianglePS(a float64, p *float64, s *float64) {
    *p = 3 * a
    *s = a*a * math.Sqrt(3) / 4
}
*/

/*
func TrianglePS(a float64) (p float64, s float64) {
    p = 3 * a
    s = a*a * math.Sqrt(3) / 4
    return
}
*/

func TrianglePS(a float64) (float64, float64) {
    p := 3 * a
    s := a*a * math.Sqrt(3) / 4
    return p, s
}

func main() {
    var a, p, s float64
    for i := 0; i < 3; i++ {
        fmt.Print("a = ")
        fmt.Scan(&a)
        //TrianglePS(a, &p, &s)
        p, s = TrianglePS(a)
        fmt.Printf("P = %.2f\tS = %.2f\n\n", p, s)
    }
}