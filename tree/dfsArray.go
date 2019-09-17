package tree

type MyStack struct {
    Items []Tree
}

// New a empty Stack
func (n *MyStack) New() *MyStack {
    n.Items = []Tree{}
    return n
}

// push the trie node to stack
func (n *MyStack) push(t Tree) {
    n.Items = append(n.Items, t)
}

// pop the last node
func (n *MyStack) pop() *Tree {
    item := n.Items[len(n.Items)-1] // return the last one -> FILO
    n.Items = n.Items[0 : len(n.Items)-1]
    return &item
}

// IsEmpty confirm if it is empty
func (n *MyStack) IsEmpty() bool {
    return len(n.Items) == 0
}

// DfsArray -> Depth-first-search 深度优先搜索
func (t *Tree) DfsArray(f func(item Item)) {
    stack := MyStack{}
    stack.New()
    stack.push(*t)

    for {
        if stack.IsEmpty() {
            break
        }
        node := stack.pop()
        near := node.Children
        for i := len(near) - 1; i >= 0; i-- {
            stack.push(*near[i])
        }
        if f != nil {
            f(node.Value)
        }
    }
}
