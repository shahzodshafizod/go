package main

import "fmt"

func main() {
    var d, m int
    fmt.Print("D = ")
    fmt.Scan(&d)
    fmt.Print("M = ")
    fmt.Scan(&m)
    switch d {
        case 1:
            switch m {
                case 1:
                    d, m = 31, 12
                default:
                    switch m {
                        case 3: d = 28
                        case 5, 7, 10, 12: d = 30
                        default: d = 31
                    }
                    m--
            }
        default: d--
    }
    fmt.Printf("D = %d\t\tM = %d\n", d, m)
}