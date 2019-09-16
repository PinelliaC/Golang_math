package tree

import "testing"

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
	//g.String()
}


func TestTree_DfsArray(t *testing.T) {
	fillTree()
	g.DfsArray()
}

func TestTree_BfsArray(t *testing.T) {
	fillTree()
	g.BfsArray()
}
