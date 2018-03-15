package stack

import "fmt"

type TStack struct {
    Top *TNode
}

func (s *TStack) Make() {
    var n, data int
    fmt.Print("How many nodes?\t")
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&data)
        s.Push(data)
    }
}

func (s *TStack) ArrayMake(array []int) {
    for index, _ := range array {
        s.Push(array[index])
    }
}

func (s TStack) Display() {
    if s.IsEmpty() { return }
    for i := s.Top; i != nil; i = i.Next {
        fmt.Print(i.Data)
        if i.Next != nil {
            fmt.Printf(" %s ", string(9472))
        }
    }
    fmt.Println(" >nil")
}

func (s *TStack) Push(d int) {
    var newNode *TNode = new(TNode)
    newNode.Data = d
    newNode.Next = s.Top
    s.Top = newNode
}

func (s *TStack) Pop() int {
    var tmpNode *TNode = s.Top
    s.Top = s.Top.Next
    return tmpNode.Data
}

func (s TStack) IsEmpty() bool {
    return s.Top == nil
}

func (s TStack) Peek() int {
    return s.Top.Data
}