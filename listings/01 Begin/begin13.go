package main

import "fmt"

func main() {
    const PI = 3.14
    var r1, r2, s1, s2, s3 float64
    fmt.Print("R1 = ")
    fmt.Scan(&r1)
    fmt.Print("R2 = ")
    fmt.Scan(&r2)
    s1 = PI * r1 * r1
    s2 = PI * r2  *r2
    s3 = s1 - s2
    fmt.Printf("S1 = %.2f\n", s1)
    fmt.Printf("S2 = %.2f\n", s2)
    fmt.Printf("S3 = %.2f\n", s3)
}