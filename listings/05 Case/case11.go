package main

import "fmt"

func main() {
    var (
        c string
        n1, n2 int
    )
    fmt.Print("Direction: C = ")
    fmt.Scan(&c)
    fmt.Print("N1 = ")
    fmt.Scan(&n1)
    fmt.Print("N2 = ")
    fmt.Scan(&n2)
    
    //1+1 = 2   180
    //1-1 = 0   continue
    //1+2 = 3   to right
    //-1+1 = 0  continue
    //-1-1 = -2 180
    //-1+2 = 1  to left
    //2+1 = 3   to right
    //2-1 = 1   to left
    //2+2 = 4   continue
    
    switch n1 + n2 {
        case 1: //to left
            switch c {
                case "N", "n": c = "W"
                case "W", "w": c = "S"
                case "S", "s": c = "E"
                case "E", "e": c = "N"
            }
        case 2, -2: //turn 180
            switch c {
                case "N", "n": c = "S"
                case "S", "s": c = "N"
                case "W", "w": c = "E"
                case "E", "e": c = "W"
            }
        case 3: //to right
            switch c {
                case "N", "n": c = "E"
                case "E", "e": c = "S"
                case "S", "s": c = "W"
                case "W", "w": c = "N"
            }
        //case 0, 4: //continue
    }
    fmt.Printf("New Direction: C = %s\n", c)
}