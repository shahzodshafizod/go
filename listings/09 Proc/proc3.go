package main

import "fmt"
import "math"

/*
func Mean(x float64, y float64, aMean *float64, gMean *float64) {
    *aMean = (x + y) / 2
    *gMean = math.Sqrt(x * y)
}
*/

/*
func Mean(x float64, y float64) (aMean float64, gMean float64) {
    aMean = (x + y) / 2
    gMean = math.Sqrt(x * y)
    return
}
*/

func Mean(x float64, y float64) (float64, float64) {
    aMean := (x + y) / 2
    gMean := math.Sqrt(x * y)
    return aMean, gMean
}

func main() {
    var a, b, c, d, aM, gM float64
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    fmt.Print("C = ")
    fmt.Scan(&c)
    fmt.Print("D = ")
    fmt.Scan(&d)
    //Mean(a, b, &aM, &gM)
    aM, gM = Mean(a, b)
    fmt.Print("Mean(A, B, AMean, GMean):\t")
    fmt.Printf("AMean = %.2f\tGMean = %.2f\n", aM, gM)
    //Mean(a, c, &aM, &gM)
    aM, gM = Mean(a, c)
    fmt.Print("Mean(A, C, AMean, GMean):\t")
    fmt.Printf("AMean = %.2f\tGMean = %.2f\n", aM, gM)
    //Mean(a, d, &aM, &gM)
    aM, gM = Mean(a, d)
    fmt.Print("Mean(A, D, AMean, GMean):\t")
    fmt.Printf("AMean = %.2f\tGMean = %.2f\n", aM, gM)
}