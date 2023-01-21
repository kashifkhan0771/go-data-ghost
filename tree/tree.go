package tree

import "errors"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

type Tree struct {
	root *TreeNode
	size int
}

// NewTree creates and returns a new tree
func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(value int) {
	if t.root == nil {
		t.root = &TreeNode{value: value}
		t.size++
	} else {
		t.root = t.insert(t.root, value)
	}
}

// Search returns true if the value is in the tree, false otherwise
func (t *Tree) Search(value int) bool {
	currentNode := t.root
	for currentNode != nil {
		if value == currentNode.value {
			return true
		}
		if value < currentNode.value {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}
	return false
}

// InOrderTraversal traverses the tree in-order and returns a slice of values
func (t *Tree) InOrderTraversal() []int {
	values := []int{}
	inOrderTraversal(t.root, &values)
	return values
}

// PreOrderTraversal traverses the tree pre-order and returns a slice of values
func (t *Tree) PreOrderTraversal() []int {
	values := []int{}
	preOrderTraversal(t.root, &values)
	return values
}

// PostOrderTraversal traverses the tree post-order and returns a slice of values
func (t *Tree) PostOrderTraversal() []int {
	values := []int{}
	postOrderTraversal(t.root, &values)
	return values
}

func (t *Tree) Delete(value int) {
	if t.root != nil {
		t.root = t.deleteNode(t.root, value)
	}
}

// Minimum finds the minimum value in tree
func (t *Tree) Minimum() int {
	node := t.root
	for node.left != nil {
		node = node.left
	}
	return node.value
}

// Maximum finds the maximum value in tree
func (t *Tree) Maximum() (int, error) {
	if t != nil {
		return 0, errors.New("tree is empty")
	}

	node := t.root
	for node.right != nil {
		node = node.right
	}

	return node.value, nil
}

//Size method
func (t *Tree) Size() int {
	return t.size
}

//Clear method
func (t *Tree) Clear() {
	t.root = nil
	t.size = 0
}
