package main

import "fmt"

func main() {
    var d, m int
    fmt.Print("D = ")
    fmt.Scan(&d)
    fmt.Print("M = ")
    fmt.Scan(&m)
    switch m {
        case 1, 3, 5, 7, 8, 10:
            switch d {
                case 31: d, m = 1, m + 1
                default: d++
            }
        case 4, 6, 9, 11:
            switch d {
                case 30: d, m = 1, m + 1
                default: d++
            }
        case 2:
            switch d {
                case 28: d, m = 1, m + 1
                default: d++
            }
        case 12:
            switch d {
                case 31: d, m = 1, 1
                default: d++
            }
    }
    fmt.Printf("D = %d\t\tM = %d\n", d, m)
}