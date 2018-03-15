package main

import (
    "fmt"
    "math"
)

func main() {
    var nomer int
    fmt.Scan(&nomer)
    switch nomer {
        case 1: fmt.Print("a = ")
        case 2: fmt.Print("c = ")
        case 3: fmt.Print("h = ")
        case 4: fmt.Print("S = ")
    }
    var value, a, c, h, s float64
    fmt.Scan(&value)
    switch nomer {
        case 1: a = value
            c = a * math.Sqrt(2);
            h = c / 2;
            s = c * h / 2;
        case 2: c = value
            a = c / math.Sqrt(2);
            h = c / 2;
            s = c * h / 2;
        case 3: h = value
            a = 2 * h / math.Sqrt(2);
            c = a * math.Sqrt(2);
            s = c * h / 2;
        case 4: s = value
            h = math.Sqrt(s);
            c = 2 * h;
            a = c / math.Sqrt(2);
    }
    if nomer != 1 {
        fmt.Printf("a = %.2f\t", a)
    }
    if nomer != 2 {
        fmt.Printf("c = %.2f\t", c)
    }
    if nomer != 3 {
        fmt.Printf("h = %.2f\t", h)
    }
    if nomer != 4 {
        fmt.Printf("S = %.2f\t", s)
    }
}