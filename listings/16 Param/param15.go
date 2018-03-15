package main

import "fmt"

func Split2(a []int, n int, b *[]int, nb *int, c *[]int, nc *int) {
    for index, _ := range a {
        if a[index] % 2 == 0 {
            *b = append(*b, a[index])
            (*nb)++
        } else {
            *c = append(*c, a[index])
            (*nc)++
        }
    }
}

func main() {
    var nA, nB, nC int
    fmt.Printf("NA = ")
    fmt.Scan(&nA)
    var a []int = make([]int, nA)
    for index, _ := range a {
        fmt.Scan(&a[index])
    }
    var b, c []int
    Split2(a, nA, &b, &nB, &c, &nC)
    fmt.Println()
    fmt.Printf("NB = %d\t", nB)
    fmt.Println(b)
    fmt.Printf("NC = %d\t", nC)
    fmt.Println(c)
}