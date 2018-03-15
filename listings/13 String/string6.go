package main

import "fmt"

func main() {
    var c string
    fmt.Print("C = ")
    fmt.Scan(&c)
    var code rune = []rune(c)[0]
    if code >= 65 && code <= 90 || code >= 97 && code <= 122 {
        fmt.Println("lat")
    } else if code >= 1040 && code <= 1103 {
        fmt.Println("rus")
    } else if code >= 48 && code <= 57 {
        fmt.Println("digit")
    }
}