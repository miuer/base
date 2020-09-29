package tree

// Node -
type Node struct {
	Value       interface{}
	FirstChild  *Node
	NextSibling *Node
}

// Tree -
type Tree struct {
	root *Node
}
