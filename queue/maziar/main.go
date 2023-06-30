package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	items []int
}

// Enqueue - add items to the end of the queue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue - remove the items from the first of the queue
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	} else {
		fIi := q.items[0]
		var fI int
		fI, q.items = q.items[fIi], q.items[1:]
		return fI, nil
	}
}

// Peek
func (q Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	} else {
		peekI := q.items[0]
		peek := q.items[peekI]
		return peek, nil
	}

}

// IsEmpty - check if the queue is empty or not
func (q Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Length - return size of the queue
func (q Queue) Length() int {
	return len(q.items)
}
func main() {
	q := Queue{}
	i, _ := q.Peek()
	fmt.Println("Peek is ", i)
	fmt.Println("is queue empty? ", q.IsEmpty())
	fmt.Println("Queue size: ", q.Length())

	q.Enqueue(1)
	fmt.Println("is queue empty? ", q.IsEmpty())
	fmt.Println("Queue size: ", q.Length())
	fmt.Println(q)
	q.Enqueue(20)
	q.Enqueue(300)
	q.Enqueue(4000)
	fmt.Println("is queue empty? ", q.IsEmpty())
	fmt.Println("Queue size: ", q.Length())
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
}
