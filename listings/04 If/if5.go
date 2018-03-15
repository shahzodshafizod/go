package main

import "fmt"

func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)
    var positives, negatives uint = 0, 0
    if a > 0 { positives++ }
    if b > 0 { positives++ }
    if c > 0 { positives++ }
    if a < 0 { negatives++ }
    if b < 0 { negatives++ }
    if c < 0 { negatives++ }
    fmt.Printf("positives = %d\nnegatives = %d\n", positives, negatives)
}