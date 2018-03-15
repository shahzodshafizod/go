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
        case 2: fmt.Print("R1 = ")
        case 3: fmt.Print("R2 = ")
        case 4: fmt.Print("S = ")
    }
    var value, a, r1, r2, s float64
    fmt.Scan(&value)
    switch nomer {
        case 1: a = value
            r1 = a * math.Sqrt(3) / 6;
            r2 = 2 * r1;
            s = a*a * math.Sqrt(3) / 4;
        case 2: r1 = value
            r2 = 2 * r1;
            a = 6 * r1 / math.Sqrt(3);
            s = a*a * math.Sqrt(3) / 4;
        case 3: r2 = value
            r1 = r2 / 2;
            a = 6 * r1 / math.Sqrt(3);
            s = a*a * math.Sqrt(3) / 4;
        case 4: s = value
            a = math.Sqrt( 4 * s / math.Sqrt(3) );
            r1 = a * math.Sqrt(3) / 6;
            r2 = 2 * r1;
    }
    if nomer != 1 {
        fmt.Printf("a = %.2f\t", a)
    }
    if nomer != 2 {
        fmt.Printf("R1 = %.2f\t", r1)
    }
    if nomer != 3 {
        fmt.Printf("R2 = %.2f\t", r2)
    }
    if nomer != 4 {
        fmt.Printf("S = %.2f\t", s)
    }
}