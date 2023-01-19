package linkedlist

func (l *LinkedList) appendSingle(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
}

func (l *LinkedList) appendDouble(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode
}

func (l *LinkedList) prependSingle(value int) {
	newNode := &Node{value: value}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) prependDouble(value int) {
	newNode := &Node{value: value}
	newNode.next = l.head
	if l.head != nil {
		l.head.prev = newNode
	}
	l.head = newNode
	if l.tail == nil {
		l.tail = newNode
	}
}

func mergeSortHelper(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	slow := head
	fast := head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	second := slow.next
	slow.next = nil
	left := mergeSortHelper(head)
	right := mergeSortHelper(second)
	return merge(left, right)
}

func merge(left *Node, right *Node) *Node {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	var head *Node
	if left.value < right.value {
		head = left
		head.next = merge(left.next, right)
	} else {
		head = right
		head.next = merge(left, right.next)
	}
	return head
}
