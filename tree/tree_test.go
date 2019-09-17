package tree

import (
    "fmt"
    "testing"
)

var g Tree

func fillTree() {
    n681 := Tree{
        Value: 681,
    }
    n587 := Tree{
        Value:    587,
        Children: []*Tree{&n681},
    }
    n162 := Tree{
        Value: 162,
    }
    n123 := Tree{
        Value:    123,
        Children: []*Tree{&n162, &n587},
    }
    n580 := Tree{
        Value: 580,
    }
    n762 := Tree{
        Value: 762,
    }
    n906 := Tree{
        Value: 906,
    }
    n945 := Tree{
        Value:    945,
        Children: []*Tree{&n580, &n762},
    }
    n131 := Tree{
        Value:    131,
        Children: []*Tree{&n906},
    }
    n879 := Tree{
        Value: 879,
    }
    g = Tree{
        Value:    110,
        Children: []*Tree{&n123, &n879, &n945, &n131},
    }
}

func TestAdd(t *testing.T) {
    fillTree()
}

func TestTree_DfsArray(t *testing.T) {
    fillTree()
    g.DfsArray(func(item Item) {
        fmt.Printf("DFS Array visiting... %v\n", item)
    })
}

func TestTree_DfsList(t *testing.T) {
    fillTree()
    g.DfsList(func(item Item) {
        fmt.Printf("DFS List visiting... %v\n", item)
    })
}

func TestTree_BfsArray(t *testing.T) {
    fillTree()
    g.BfsList(func(item Item) {
        fmt.Printf("BFS Array visiting... %v\n", item)
    })
}

func TestTree_BfsList(t *testing.T) {
    fillTree()
    g.BfsList(func(item Item) {
        fmt.Printf("BFS LIST visiting... %v\n", item)
    })
}

func BenchmarkTree_BfsList(b *testing.B) {
    fillTree()
    g.BfsList(nil)
}

func BenchmarkTree_BfsArray(b *testing.B) {
    fillTree()
    g.BfsArray(nil)
}

func BenchmarkTree_DfsArray(b *testing.B) {
    fillTree()
    g.DfsArray(nil)
}

func BenchmarkTree_DfsList(b *testing.B) {
    fillTree()
    g.DfsList(nil)
}
