package main

import (
    "fmt"
    "math"
)

func main() {
    const PI = 3.14
    var nomer int
    fmt.Scan(&nomer)
    switch nomer {
        case 1: fmt.Print("R = ")
        case 2: fmt.Print("D = ")
        case 3: fmt.Print("L = ")
        case 4: fmt.Print("S = ")
    }
    var value, r, d, l, s float64
    fmt.Scan(&value)
    switch nomer {
        case 1: r = value
            d = 2 * r;
            l = 2 * PI * r;
            s = PI * r * r;
        case 2: d = value
            r = d / 2;
            l = 2 * PI * r;
            s = PI * r * r;
        case 3: l = value
            r = l / (2 * PI);
            d = 2 * r;
            s = PI * r * r;
        case 4: s = value
            r = math.Sqrt(s / PI);
            d = 2 * r;
            l = 2 * PI * r;
    }
    if nomer != 1 {
        fmt.Printf("R = %.2f\t", r)
    }
    if nomer != 2 {
        fmt.Printf("D = %.2f\t", d)
    }
    if nomer != 3 {
        fmt.Printf("L = %.2f\t", l)
    }
    if nomer != 4 {
        fmt.Printf("S = %.2f\t", s)
    }
}