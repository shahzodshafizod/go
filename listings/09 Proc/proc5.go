package main

import (
    "fmt"
    "math"
)

/*
func RectPS(x1 float64, y1 float64, x2 float64, y2 float64, p *float64, s *float64) {
    a := math.Abs(x2 - x1)
    b := math.Abs(y2 - y1)
    *p = 2 * (a + b)
    *s = a * b
}
*/

/*
func RectPS(x1 float64, y1 float64, x2 float64, y2 float64) (p float64, s float64) {
    a := math.Abs(x2 - x1)
    b := math.Abs(y2 - y1)
    p = 2 * (a + b)
    s = a * b
    return
}
*/

func RectPS(x1 float64, y1 float64, x2 float64, y2 float64) (float64, float64) {
    a := math.Abs(x2 - x1)
    b := math.Abs(y2 - y1)
    p := 2 * (a + b)
    s := a * b
    return p, s
}

func main() {
    var x1, y1, x2, y2, p, s float64
    for i := 0; i < 3; i++ {
        fmt.Scan(&x1, &y1, &x2, &y2)
        //RectPS(x1, y1, x2, y2, &p, &s)
        p, s = RectPS(x1, y1, x2, y2)
        fmt.Printf("P = %.2f\t\tS = %.2f\n\n", p, s)
    }
}