package dsa

import "github.com/pkg/errors"

var (
	errStackFull  = errors.New("stack is full")
	errStackEmpty = errors.New("stack is empty")
)

// 顺序队列
type SqStack struct {
	data []int
	top  int
}

func NewSqStack(maxSize int) *SqStack {
	return &SqStack{
		data: make([]int, maxSize),
		top:  -1,
	}
}

// 判空
func (s *SqStack) IsEmpty() bool {
	return s.top == -1
}

// 判满
func (s *SqStack) IsFull() bool {
	return s.top+1 == len(s.data)
}

// 进栈
func (s *SqStack) Push(data int) error {
	if s.IsFull() {
		return errStackFull
	}
	s.top++
	s.data[s.top] = data
	return nil
}

// 出栈
func (s *SqStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errStackEmpty
	}
	data := s.data[s.top]
	s.top--
	return data, nil
}

// TODO: 链表队列
type LStack LNode
