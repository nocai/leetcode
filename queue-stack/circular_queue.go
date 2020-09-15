package leetcode

type CircularQueue struct {
	data []int
	head int
	tail int
}

func NewCircularQueue(k int) *CircularQueue {
	return &CircularQueue{
		data: make([]int, 0, k),
		head: -1,
		tail: -1,
	}
}

func (q *CircularQueue) enQueue(value int) bool {
	if q.isFull() {
		return false
	}

	if q.isEmpty() {
		q.head = 0
	}

	q.tail = (q.tail + 1) % len(q.data)
	q.data[q.tail] = value
	return true
}

func (q *CircularQueue) deQueue() bool {
	if q.isEmpty() {
		return false
	}

	if q.head == q.tail {
		q.head = -1
		q.tail = -1
		return true
	}

	q.head = (q.head + 1) % len(q.data)
	return true
}

func (q *CircularQueue) front() int {
	if q.isEmpty() {
		return -1
	}
	return q.data[q.head]
}

func (q *CircularQueue) rear() int {
	if q.isEmpty() {
		return -1
	}
	return q.data[q.tail]
}

func (q *CircularQueue) isEmpty() bool {
	return q.head == -1
}

func (q *CircularQueue) isFull() bool {
	return (q.tail+1)%len(q.data) == q.head
}
