package tree

type TNode struct {
    Data int
    Left *TNode
    Right *TNode
    Parent *TNode
}