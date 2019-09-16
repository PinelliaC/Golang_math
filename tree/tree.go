package tree

type Item interface {
}
type Tree struct {
	Value    Item
	Children []*Tree
}
