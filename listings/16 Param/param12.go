package main

import "fmt"

func SortIndex(a []float32, n int, i *[]int) {
    *i = make([]int, n)
    for index, _ := range *i {
        (*i)[index] = index
    }
    for k := 0; k < n-1; k++ {
        for j := 1; j < n-k; j++ {
            if a[(*i)[j-1]] > a[(*i)[j]] {
                tmp := (*i)[j-1]
                (*i)[j-1] = (*i)[j]
                (*i)[j] = tmp
            }
        }
    }
    for index, _ := range *i {
        (*i)[index]++
    }
}

func main() {
    var n int
    for i := 1; i <= 3; i++ {
        fmt.Printf("N%d = ", i)
        fmt.Scan(&n)
        var array []float32 = make([]float32, n)
        for index, _ := range array {
            fmt.Scan(&array[index])
        }
        var indexs []int
        SortIndex(array, n, &indexs)
        fmt.Println(indexs)
    }
}