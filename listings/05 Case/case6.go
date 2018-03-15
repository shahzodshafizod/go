package main

import "fmt"

func main() {
	var nomer int
	fmt.Scan(&nomer)
    switch nomer {
        case 1: fmt.Print("Length (in dm):\t")
        case 2: fmt.Print("Length (in km):\t")
        case 3: fmt.Print("Length (in m):\t")
        case 4: fmt.Print("Length (in mm):\t")
        case 5: fmt.Print("Length (in sm):\t")
	}
    var value float64
    fmt.Scan(&value)
    switch nomer {
        case 1: value /= 10
        case 2: value *= 1000
        //case 3:
        case 4: value /= 1000
        case 5: value /= 100
    }
    fmt.Println("Length (in m):\t", value)
}
