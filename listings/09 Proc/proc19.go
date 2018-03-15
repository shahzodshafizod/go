package main

import "fmt"

func RingS(r1 float32, r2 float32) float32 {
    const PI = 3.14
    s1 := PI * r1 * r1
    s2 := PI * r2 * r2
    return s1 - s2
}

func main() {
    var r1, r2, s float32
    for i := 1; i <= 3; i++ {
        fmt.Printf("R1 = ")
        fmt.Scan(&r1)
        fmt.Printf("R2 = ")
        fmt.Scan(&r2)
        s = RingS(r1, r2)
        fmt.Printf("RingS(%.2f, %.2f) = %.2f\n\n", r1, r2, s)
    }
}