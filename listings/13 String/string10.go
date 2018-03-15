package main

import "fmt"

func reverse(str string) string {
    var codes []rune = []rune(str)
    from, to := 0, len(codes) - 1
    var tmp rune
    for from < to {
        tmp = codes[from]
        codes[from] = codes[to]
        codes[to] = tmp
        from++
        to--
    }
    return string(codes)
}

func main() {
    var str string
    fmt.Print("string:\t")
    fmt.Scan(&str)
    str = reverse(str)
    fmt.Println(str)
}