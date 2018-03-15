package main

import "fmt"

func int2string(number int) string {
    isNegative := false
    if number < 0 {
        isNegative = true
        number *= -1
    }
    start := []rune("0")[0]
    var array []rune
    for number > 0 {
        array = append(array, start + rune(number % 10))
        number /= 10
    }
    if isNegative {
        array = append(array, []rune("-")[0])
    }
    from, to := 0, len(array) - 1
    for from < to {
        tmp := array[from]
        array[from] = array[to]
        array[to ] = tmp
        from++
        to--
    }
    return string(array)
}

func main() {
    var number int
    fmt.Scanf("%d", &number)
    var str string = int2string(number)
    darozi := len(str)
    for index := 0; index < darozi; index++ {
        fmt.Printf("%c\t", str[index])
    }
}