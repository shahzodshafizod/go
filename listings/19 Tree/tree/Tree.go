package tree

import "fmt"
import "math/rand"

type Tree struct {
    root *TNode
    current *TNode
    field *Field
    level int
    nodeCount int
}

func (t *Tree) Make() {
    if t.root == nil {
        t.root = new(TNode)
        t.nodeCount++
        fmt.Print("root's Data: ")
        fmt.Scan(&(t.root.Data))
        t.current = t.root
    }
    var answer int
    fmt.Print("Where to go? (0-Exit; 1-Left; 2-Right; 3-Parent):\t")
    fmt.Scan(&answer)
    if answer != 1 && answer != 2 {
        switch answer {
            case 0: t.current = t.root; return
            case 3: t.current = t.current.Parent;
        }
        t.Make()
        return
    }
    var newNode *TNode = new(TNode)
    t.nodeCount++
    fmt.Print("Data = ")
    fmt.Scan(&(newNode.Data))
    newNode.Parent = t.current
    if answer == 1 {
        t.current.Left = newNode
    } else {
        t.current.Right = newNode
    }
    t.current = newNode
    t.Make()
}

func (t *Tree) AutoMake(count int) {
    if count <= 0 {
        t.current = t.root
        return
    }
    if t.root == nil {
        t.root = new(TNode)
        t.root.Data = rand.Intn(90) + 10
        t.nodeCount++
        count--
        t.current = t.root
        if count <= 0 { return }
    }
    var direction int
    for {
        direction = rand.Intn(3) + 1
        if !(direction == 3 && t.current.Parent == nil) {
            break
        }
    }
    var exit = false
    switch direction {
        case 1:
            if t.current.Left != nil {
                t.current = t.current.Left
                exit = true
            }
        case 2:
            if t.current.Right != nil {
                t.current = t.current.Right
                exit = true
            }
        case 3:
            t.current = t.current.Parent
            exit = true
    }
    if exit {
        t.AutoMake(count)
        return
    }
    var newNode *TNode = new(TNode)
    t.nodeCount++
    count--
    newNode.Data = rand.Intn(90) + 10
    newNode.Parent = t.current
    if direction == 1 {
        t.current.Left = newNode
    } else {
        t.current.Right = newNode
    }
    t.current = newNode
    t.AutoMake(count)
}

func (t Tree) GetNodeCount() int {
    return t.nodeCount
}

func (t* Tree) GetLevel() int {
    if t.level == 0 {
        t.setLevel(0)
    }
    return t.level;
}

func (t* Tree) setLevel(level int) {
    if (t.current == nil) {
        t.current = t.root
        return
    }
    if level > t.level {
        t.level = level
    }
    var current *TNode = t.current
    t.current = current.Left
    t.setLevel(level + 1)
    t.current = current.Right
    t.setLevel(level + 1)
}

func (t Tree) Display() {
    if t.root == nil { return }
    if t.field == nil {
        t.field = new(Field)
        t.field.Init(t.root, t.GetLevel()+1, t.GetNodeCount())
    }
    t.field.Display()
}

func (t Tree) GetDataCount(data int) int {
    if t.current == nil {
        t.current = t.root
        return 0
    }
    var current *TNode = t.current
    var count = 0
    if t.current.Data == data {
        count++
    }
    t.current = current.Left
    count += t.GetDataCount(data)
    t.current = current.Right
    count += t.GetDataCount(data)
    return count
}

func (t Tree) GetDataSum() int {
    if t.current == nil {
        t.current = t.root
        return 0
    }
    var current *TNode = t.current
    sum := current.Data
    t.current = current.Left
    sum += t.GetDataSum()
    t.current = current.Right
    sum += t.GetDataSum()
    return sum
}

func (t Tree) GetLeftCount() int {
    if t.current == nil {
        t.current = t.root
        return 0
    }
    var current *TNode = t.current
    var count = 0
    if current.Parent != nil && current.Parent.Left == current {
        count++
    }
    t.current = current.Left
    count += t.GetLeftCount()
    t.current = current.Right
    count += t.GetLeftCount()
    return count
}

func (t Tree) GetLeafCount() int {
    if t.current == nil {
        t.current = t.root
        return 0
    }
    var current *TNode = t.current
    var count = 0
    if current.Left == nil && current.Right == nil {
        count++
    }
    t.current = current.Left
    count += t.GetLeafCount()
    t.current = current.Right
    count += t.GetLeafCount()
    return count
}

func (t Tree) GetLeafSum() int {
    if t.current == nil {
        t.current = t.root
        return 0
    }
    var current *TNode = t.current
    var sum = 0
    if current.Left == nil && current.Right == nil {
        sum += current.Data
    }
    t.current = current.Left
    sum += t.GetLeafSum()
    t.current = current.Right
    sum += t.GetLeafSum()
    return sum
}

func (t Tree) GetRightLeafCount() int {
    if t.current == nil {
        t.current = t.root
        return 0
    }
    var current *TNode = t.current
    var count int = 0
    if current.Left == nil && current.Right == nil && 
        current.Parent != nil && current.Parent.Right == current {
            count++
    }
    t.current = current.Left
    count += t.GetRightLeafCount()
    t.current = current.Right
    count += t.GetRightLeafCount()
    return count
}

func (t Tree) LevelToArray(array []int, level int, toCount bool) {
    if t.current == nil {
        t.current = t.root
        return
    }
    var current *TNode = t.current
    if toCount {
        array[level]++
    } else {
        array[level] += current.Data
    }
    t.current = current.Left
    t.LevelToArray(array, level+1, toCount)
    t.current = current.Right
    t.LevelToArray(array, level+1, toCount)
}

func (t Tree) Infix() {
    if t.current == nil {
        t.current = t.root
        return
    }
    var current *TNode = t.current
    t.current = current.Left
    t.Infix()
    fmt.Printf("%d\t", current.Data)
    t.current = current.Right
    t.Infix()
}

func (t Tree) Prefix() {
    if t.current == nil {
        t.current = t.root
        return
    }
    var current *TNode = t.current
    fmt.Printf("%d\t", current.Data)
    t.current = current.Left
    t.Prefix()
    t.current = current.Right
    t.Prefix()
}

func (t Tree) Postfix() {
    if t.current == nil {
        t.current = t.root
        return
    }
    var current *TNode = t.current
    t.current = current.Left
    t.Postfix()
    t.current = current.Right
    t.Postfix()
    fmt.Printf("%d\t", current.Data)
}

func (t Tree) InfixToN(index *int, n int) {
    if t.current == nil {
        t.current = t.root
        return
    }
    var current *TNode = t.current
    t.current = current.Left
    t.InfixToN(index, n)
    if *index < n {
        fmt.Printf("%d\t", current.Data)
        (*index)++
    } else {
        t.current = t.root
        return
    }
    t.current = current.Right
    t.InfixToN(index, n)
}

func (t Tree) PostfixFromN(index *int, n int) {
    if t.current == nil {
        t.current = t.root
        return
    }
    var current *TNode = t.current
    t.current = current.Left
    t.PostfixFromN(index, n)
    t.current = current.Right
    t.PostfixFromN(index, n)
    (*index)++
    if *index >= n {
        fmt.Printf("%d\t", current.Data)
    }
}

func (t Tree) PrefixBetween(index *int, n1 int, n2 int) {
    if t.current == nil {
        t.current = t.root
        return
    }
    (*index)++
    if *index >= n1 && *index <= n2 {
        fmt.Printf("%d\t", t.current.Data)
    } else if *index > n2 {
        t.current = t.root
        return
    }
    var current *TNode = t.current
    t.current = current.Left
    t.PrefixBetween(index, n1, n2)
    t.current = current.Right
    t.PrefixBetween(index, n1, n2)
}

func (t Tree) PrintLevel(index int, level int) int {
    if t.current == nil {
        t.current = t.root
        return 0
    }
    var count = 0
    if index == level {
        count++
        fmt.Printf("%d\t", t.current.Data)
    }
    var current *TNode = t.current
    t.current = current.Left
    count += t.PrintLevel(index+1, level)
    t.current = current.Right
    count += t.PrintLevel(index+1, level)
    return count
}

func (t Tree) GetMaxData() int {
    if t.current == nil {
        t.current = t.root
        return 0
    }
    var current *TNode = t.current
    var maximal int = current.Data
    if current.Left != nil {
        t.current = current.Left
        data := t.GetMaxData()
        if data > maximal {
            maximal = data
        }
    }
    if current.Right != nil {
        t.current = current.Right
        data := t.GetMaxData()
        if data > maximal {
            maximal = data
        }
    }
    t.current = t.root
    return maximal
}

func (t Tree) GetMinData() int {
    if t.current == nil {
        t.current = t.root
        return 0
    }
    var current *TNode = t.current
    var minimal int = current.Data
    if current.Left != nil {
        t.current = current.Left
        data := t.GetMinData()
        if data < minimal {
            minimal = data
        }
    }
    if current.Right != nil {
        t.current = current.Right
        data := t.GetMinData()
        if data < minimal {
            minimal = data
        }
    }
    t.current = t.root
    return minimal
}

func (t Tree) GetLeafDataCount(data int) int {
    if t.current == nil {
        t.current = t.root
        return 0
    }
    var current *TNode = t.current
    var count = 0
    if current.Data == data && current.Left == nil && current.Right == nil {
        count++
    }
    t.current = current.Left
    count += t.GetLeafDataCount(data)
    t.current = current.Right
    count += t.GetLeafDataCount(data)
    return count
}