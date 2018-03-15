package main

import "fmt"

func main() {
    var K int
    fmt.Print("K = ")
    fmt.Scanf("%d", &K)
    mark := "";
    switch K {
        case 1: mark = "bad"
        case 2: mark = "ghayriqanoatbaxsh"
        case 3: mark = "qanoatbaxsh"
        case 4: mark = "xub"
        case 5: mark = "a'lo"
        default: mark = "error"
    }
    fmt.Println("mark:", mark)
}