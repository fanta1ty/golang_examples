package main

import "fmt"

type Queue struct {
	queue []int
}

func newQueue() Queue {
	return Queue{queue: []int{}}
}

func (q *Queue) enQueue(item int) {
	q.queue = append(q.queue, item)
}

func (q *Queue) deQueue() (int, error) {
	if len(q.queue) <= 0 {
		return 0, fmt.Errorf("queue is empty")
	} else if len(q.queue) == 1 {
		item := q.queue[0]
		q.queue = []int{}
		return item, nil
	}

	item := q.queue[0]
	q.queue = q.queue[1:]
	return item, nil
}

func main() {
	q := newQueue()
	q.enQueue(1)
	q.enQueue(2)
	fmt.Println(q.deQueue())
	fmt.Println(q.deQueue())
}
