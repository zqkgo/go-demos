package dsa

import (
	"github.com/pkg/errors"
)

var (
	errQueueEmpty = errors.New("queue is empty")
	errQueueFull  = errors.New("queue is full")
)

// 顺序队列，循环队列
type SqQueue struct {
	data  []int
	front int
	rear  int
}

func NewSqQueue(maxSize int) *SqQueue {
	return &SqQueue{
		data: make([]int, maxSize),
	}
}

// 判空
func (q *SqQueue) IsEmpty() bool {
	return q.front == q.rear
}

// 判满
func (q *SqQueue) IsFull() bool {
	return (q.rear+1)%len(q.data) == q.front
}

// 入队
func (q *SqQueue) EnQueue(data int) error {
	if q.IsFull() {
		return errQueueFull
	}
	q.rear = (q.rear + 1) % len(q.data)
	q.data[q.rear] = data
	return nil
}

// 出队
func (q *SqQueue) DeQueue() (int, error) {
	if q.IsEmpty() {
		return 0, errQueueEmpty
	}
	q.front = (q.front + 1) % len(q.data)
	return q.data[q.front], nil
}

// TODO: 链表队列
type LQueue struct {

}