package main

import "fmt"

func main() {
    var x, y, z int
    fmt.Print("number1 = ")
    fmt.Scan(&x)
    fmt.Print("number2 = ")
    fmt.Scan(&y)
    fmt.Print("number3 = ")
    fmt.Scan(&z)
    positives := 0
    if x > 0 { positives++ }
    if y > 0 { positives++ }
    if z > 0 { positives++ }
    fmt.Printf("positives: %d\n", positives)
}