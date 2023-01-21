package queue

type Queue struct {
	items []int
}

// NewQueue creates and returns a new queue
func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue adds a value to the end of the queue
func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

// Dequeue removes and returns the front value of the queue
func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	val := q.items[0]
	q.items = q.items[1:]
	return val, true
}

// Peek returns the front value of the queue
func (q *Queue) Peek() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	return q.items[0], true
}

// Size returns the size of the queue
func (q *Queue) Size() int {
	return len(q.items)
}
