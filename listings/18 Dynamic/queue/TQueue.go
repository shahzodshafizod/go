package queue

import "fmt"

type TQueue struct {
    Head *TNode
    Tail *TNode
}

func (s *TQueue) Make() {
    var n, data int
    fmt.Print("How many nodes?\t")
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&data)
        s.Enqueue(data)
    }
}

func (s TQueue) Display() {
    if s.IsEmpty() { return }
    for i := s.Head; i != nil; i = i.Next {
        fmt.Print(i.Data)
        if i.Next != nil {
            fmt.Printf(" %s ", string(9472))
        }
    }
    fmt.Println(" >nil")
}

func (s *TQueue) Enqueue(d int) {
    var newNode *TNode = new(TNode)
    newNode.Data = d
    if s.Head == nil {
        s.Head, s.Tail = newNode, newNode
    } else {
        s.Tail.Next = newNode
        s.Tail = newNode
    }
}

func (s *TQueue) Dequeue() int {
    var tmpNode = s.Head
    s.Head = s.Head.Next
    if s.Head == nil {
        s.Tail = nil
    }
    return tmpNode.Data
}

func (s TQueue) IsEmpty() bool {
    return s.Head == nil
}

func (s TQueue) First() int {
    return s.Head.Data
}

func (s TQueue) Last() int {
    return s.Tail.Data
}