package leetcode

type Queue struct {
	data  []int
	start int
}

func NewQueue() *Queue {
	data := make([]int, 0, 10)
	return &Queue{
		data:  data,
		start: 0,
	}
}

func (q *Queue) enQueue(x int) bool {
	q.data = append(q.data, x)
	return true
}

func (q *Queue) deQueue() bool {
	if q.isEmpty() {
		return false
	}

	q.start++
	return true
}

func (q *Queue) Front() int {
	return q.data[q.start]
}

func (q *Queue) isEmpty() bool {
	return q.start >= len(q.data)
}
