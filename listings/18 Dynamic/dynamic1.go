package main

import (
    "fmt"
    "./stack"
)

func main() {
    var P1 *stack.TNode = new(stack.TNode)
    fmt.Print("P1.Data = ")
    fmt.Scan(&(P1.Data))
    
    P1.Next = new(stack.TNode)
    fmt.Print("P1.Next.Data = ")
    fmt.Scan(&(P1.Next.Data))
    
    var P2 *stack.TNode = P1.Next
    fmt.Printf("\nP1.Data = %d\n", P1.Data)
    fmt.Printf("P2.Data = %d\n", P2.Data)
    fmt.Printf("P2 = %p", P2)
}