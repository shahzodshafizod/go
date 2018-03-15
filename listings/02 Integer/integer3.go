package main

import "fmt"

func main() {
    var bytes int
    fmt.Print("bytes = ")
    fmt.Scanf("%d", &bytes)
    kBytes := bytes / 1024
    fmt.Printf("kBytes = %d\n", kBytes)
}