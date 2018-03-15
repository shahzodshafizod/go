package main

import "./tree"
import "fmt"

func main() {
    var P1 *tree.TNode = new(tree.TNode)
    fmt.Print("P1.Data = ")
    fmt.Scan(&(P1.Data))
    
    P1.Left = new(tree.TNode)
    fmt.Print("P1.Left.Data = ")
    fmt.Scan(&(P1.Left.Data))
    
    P1.Right = new(tree.TNode)
    fmt.Print("P1.Right.Data = ")
    fmt.Scan(&(P1.Right.Data))
    
    fmt.Printf("P1.Data = %d\n", P1.Data)
    fmt.Printf("P1.Left.Data = %d\n", P1.Left.Data)
    fmt.Printf("P1.Right.Data = %d\n", P1.Right.Data)
    fmt.Printf("P2 = %p\n", P1.Left)
    fmt.Printf("P3 = %p\n", P1.Right)
}