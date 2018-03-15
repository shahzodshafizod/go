package main

import "fmt"

func main() {
	var nomer int
	fmt.Scan(&nomer)
    switch nomer {
        case 1: fmt.Print("Weight (in kg):\t")
        case 2: fmt.Print("Weight (in mg):\t")
        case 3: fmt.Print("Weight (in g):\t")
        case 4: fmt.Print("Weight (in tn):\t")
        case 5: fmt.Print("Weight (in ct):\t")
	}
    var value float64
    fmt.Scan(&value)
    switch nomer {
        //case 1: 
        case 2: value /= 1000000
        case 3: value /= 1000
        case 4: value *= 1000
        case 5: value *= 100
    }
    fmt.Println("Weight (in kg):\t", value)
}
