package main

import "fmt"

func main() {
    var (
        c string
        n int
    )
    fmt.Print("Direction: C = ")
    fmt.Scan(&c)
    fmt.Print("N = ")
    fmt.Scan(&n)
    switch c {
        case "N", "n":
            switch n {
                case 1: c = "W"
                case -1: c = "E"
            }
        case "S", "s":
            switch n {
                case 1: c = "E"
                case -1: c = "W"
            }
        case "W", "w":
            switch n {
                case 1: c = "S"
                case -1: c = "N"
            }
        case "E", "e":
            switch n {
                case 1: c = "N"
                case -1: c = "S"
            }
    }
    fmt.Printf("New Direction: C = %s\n", c)
}