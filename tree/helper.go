package tree

func (t *Tree) insert(node *TreeNode, value int) *TreeNode {
	if node == nil {
		t.size++
		return &TreeNode{value: value}
	}
	if value < node.value {
		node.left = t.insert(node.left, value)
	} else {
		node.right = t.insert(node.right, value)
	}
	return node
}

func inOrderTraversal(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}
	inOrderTraversal(node.left, values)
	*values = append(*values, node.value)
	inOrderTraversal(node.right, values)
}

func preOrderTraversal(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}
	*values = append(*values, node.value)
	preOrderTraversal(node.left, values)
	preOrderTraversal(node.right, values)
}

func postOrderTraversal(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}
	postOrderTraversal(node.left, values)
	postOrderTraversal(node.right, values)
	*values = append(*values, node.value)
}

func (t *Tree) deleteNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil
	}
	if value < node.value {
		node.left = t.deleteNode(node.left, value)
	} else if value > node.value {
		node.right = t.deleteNode(node.right, value)
	} else {
		if node.left == nil && node.right == nil {
			node = nil
			t.size--
			return nil
		}
		if node.left == nil {
			t.size--
			return node.right
		}
		if node.right == nil {
			t.size--
			return node.left
		}
		smallestValue := findSmallestValue(node.right)
		node.value = smallestValue
		node.right = t.deleteNode(node.right, smallestValue)
	}
	return node
}

func findSmallestValue(node *TreeNode) int {
	for node.left != nil {
		node = node.left
	}
	return node.value
}
