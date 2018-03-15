package main

import "fmt"

func Split1(a []float32, n int, b *[]float32, nb *int, c *[]float32, nc *int) {
    *nb = n / 2 + n % 2
    *nc = n / 2
    *b = make([]float32, *nb)
    *c = make([]float32, *nc)
    bIndex, cIndex := 0, 0
    for index, _ := range a {
        if index % 2 == 0 {
            (*b)[bIndex] = a[index]
            bIndex++
        } else {
            (*c)[cIndex] = a[index]
            cIndex++
        }
    }
}

func printArray(a []float32, n int) {
    for index, _ := range a {
        fmt.Printf("%.2f\t", a[index])
    }
    fmt.Println()
}

func main() {
    var nA, nB, nC int
    fmt.Printf("NA = ")
    fmt.Scan(&nA)
    var a []float32 = make([]float32, nA)
    for index, _ := range a {
        fmt.Scan(&a[index])
    }
    var b, c []float32
    Split1(a, nA, &b, &nB, &c, &nC)
    fmt.Println()
    fmt.Printf("NB = %d\t", nB)
    printArray(b, nB)
    fmt.Printf("NC = %d\t", nC)
    printArray(c, nC)
}