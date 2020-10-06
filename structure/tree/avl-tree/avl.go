package tree

// Node -
type Node struct {
	height int
	index  int
	Value  interface{}

	left  *Node
	right *Node
}

// AVLTree -
type AVLTree struct {
	root *Node
}

// Init -
func (avl *AVLTree) Init() *AVLTree {
	avl.root = nil
	return avl
}

// NewAVLTree -
func NewAVLTree() *AVLTree {
	return new(AVLTree).Init()
}

// Height -
func (node *Node) Height() int {
	if node == nil {
		return -1
	}

	return node.height
}

// left left
// α -> right
// α.left -> root
func (avl *AVLTree) rotateWithLeftChild(k *Node) *Node {
	n := k.left
	k.left = n.right
	n.right = k

	k.height = max(k.left.Height(), k.right.Height()) + 1
	n.height = max(n.left.Height(), k.height) + 1

	return n
}

// right right
// α -> left
// α.right -> root
func (avl *AVLTree) rotateWithRightChild(k *Node) *Node {
	n := k.right
	k.right = n.left
	n.left = k

	k.height = max(k.left.Height(), k.right.Height()) + 1
	n.height = max(n.right.Height(), k.height) + 1

	return n
}

// left right
// α -> right
// hightest -> root
// mid -> left

// α.left ->rotateWithRightChild
// α -> rotateWithLeftChild
func (avl *AVLTree) doubleWithLeftChild(k *Node) {
	k.left = avl.rotateWithRightChild(k.left)
	avl.rotateWithLeftChild(k)
}

// right left
// α -> left
// hightest -> root
// mid -> right

// α.right ->rotateWithLeftChild
// α -> rotateWithRightChild
func (avl *AVLTree) doubleWithRightChild(k *Node) {
	k.right = avl.rotateWithLeftChild(k.right)
	avl.rotateWithRightChild(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// Insert -
func (avl *AVLTree) Insert() {
}

func (avl *AVLTree) insert() {

}
