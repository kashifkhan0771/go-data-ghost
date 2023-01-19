package linkedlist

import "errors"

// Node represents a node in the linked list
type Node struct {
	value int
	next  *Node
	prev  *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	head     *Node
	tail     *Node
	isDouble bool
}

// NewLinkedList creates and returns a new linked list
func NewLinkedList(isDouble bool) *LinkedList {
	return &LinkedList{isDouble: isDouble}
}

// Append appends a value to the end of the list
func (l *LinkedList) Append(value int) {
	if l.isDouble {
		l.appendDouble(value)
	} else {
		l.appendSingle(value)
	}
}

// Prepend prepends a value to the beginning of the list
func (l *LinkedList) Prepend(value int) {
	if l.isDouble {
		l.prependDouble(value)
	} else {
		l.prependSingle(value)
	}
}

// InsertAfter insert a value after a specific node
func (l *LinkedList) InsertAfter(node *Node, value int) error {
	if !l.isDouble {
		return errors.New("InsertAfter is not supported for singly linked lists")
	}
	newNode := &Node{value: value}
	newNode.prev = node
	newNode.next = node.next
	if node.next != nil {
		node.next.prev = newNode
	} else {
		l.tail = newNode
	}
	node.next = newNode
	return nil
}

// DeleteNode delete a specific node
func (l *LinkedList) DeleteNode(node *Node) {
	if !l.isDouble {
		if node == l.head {
			l.head = node.next
		}
		currentNode := l.head
		for currentNode.next != nil {
			if currentNode.next == node {
				currentNode.next = node.next
				break
			}
			currentNode = currentNode.next
		}
	} else {
		if node == l.head {
			l.head = node.next
		}
		if node == l.tail {
			l.tail = node.prev
		}
		if node.prev != nil {
			node.prev.next = node.next
		}
		if node.next != nil {
			node.next.prev = node.prev
		}
	}
}

//Values traverse the list and return all the values in the list
func (l *LinkedList) Values() []int {
	values := make([]int, 0)
	currentNode := l.head
	for currentNode != nil {
		values = append(values, currentNode.value)
		currentNode = currentNode.next
	}
	return values
}

//Size return the size of the list
func (l *LinkedList) Size() int {
	size := 0
	currentNode := l.head
	for currentNode != nil {
		size++
		currentNode = currentNode.next
	}
	return size
}

//Find a specific node in the list based on its value
func (l *LinkedList) Find(value int) *Node {
	currentNode := l.head
	for currentNode != nil {
		if currentNode.value == value {
			return currentNode
		}
		currentNode = currentNode.next
	}
	return nil
}

// Reverse the linked list
func (l *LinkedList) Reverse() error {
	if !l.isDouble {
		return errors.New("reverse is not supported for singly linked lists")
	}
	currentNode := l.head
	var prev *Node
	l.tail = l.head
	for currentNode != nil {
		next := currentNode.next
		currentNode.next = prev
		prev = currentNode
		currentNode = next
	}
	l.head = prev

	return nil
}

// Merge two linked lists
func (l *LinkedList) Merge(l2 *LinkedList) error {
	if !l.isDouble {
		return errors.New("merge is not supported for singly linked lists")
	}

	l.tail.next = l2.head
	l2.head.prev = l.tail
	l.tail = l2.tail

	return nil
}

// BubbleSort the linked list
func (l *LinkedList) BubbleSort() error {
	if !l.isDouble {
		return errors.New("BubbleSort is not supported for singly linked lists")
	}

	var swapped bool
	currentNode := l.head
	for currentNode != nil {
		if currentNode.next != nil && currentNode.value > currentNode.next.value {
			currentNode.value, currentNode.next.value = currentNode.next.value, currentNode.value
			swapped = true
		}
		currentNode = currentNode.next
	}
	if swapped {
		if err := l.BubbleSort(); err != nil {
			return err
		}
	}

	return nil
}

//MergeSort the linked list
func (l *LinkedList) MergeSort() error {
	if !l.isDouble {
		return errors.New("MergeSort is not supported for singly linked lists")
	}
	l.head = mergeSortHelper(l.head)
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	l.tail = currentNode

	return nil
}
