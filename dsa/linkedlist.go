package dsa

import (
	"github.com/pkg/errors"
)

type LNode struct {
	Data int
	Next *LNode
}

func NewLNode(data int) *LNode {
	return &LNode{
		Data: data,
	}
}

// 头插法
func (l *LNode) InsertAtHead(data ...int) {
	for i := 0; i < len(data); i++ {
		n := NewLNode(data[i])
		n.Next = l.Next
		l.Next = n
	}
}

// 尾插法
func (l *LNode) InsertAtTail(data ...int) {
	list := l
	for i := 0; i < len(data); i++ {
		for list.Next != nil {
			list = list.Next
		}
		list.Next = NewLNode(data[i])
	}
}

// 查找
func (l *LNode) FindAt(n int) (int, error) {
	if n <= 0 {
		return -1, errors.New("n too small")
	}
	if l.IsEmpty() {
		return -1, errors.New("list is empty")
	}
	idx := 0
	list := l
	for list.Next != nil && idx < n {
		list = list.Next
		idx++
	}
	if idx != n {
		return -1, errors.New("n too big")
	}
	return list.Data, nil
}

// 删除
func (l *LNode) DeleteAt(n int) (error) {
	if n <= 0 {
		return errors.New("n too small")
	}
	idx := 0
	list := l
	var prev *LNode
	for list.Next != nil && idx < n {
		prev = list
		list = list.Next
		idx++
	}
	if idx != n {
		return errors.New("n too big")
	}
	prev.Next = list.Next
	return nil
}

// 是否为空
func (l *LNode) IsEmpty() bool {
	return l.Next == nil
}

// 重置
func (l *LNode) Reset() {
	l.Next = nil
}

func (l *LNode) Print() []int {
	list := l
	var nums []int
	for list.Next != nil {
		nums = append(nums, list.Next.Data)
		list = list.Next
	}
	return nums
}

type DLNode struct {
	Data  int
	Prior *DLNode
	Next  *DLNode
}

func NewDLNode(data int) *DLNode {
	return &DLNode{
		Data: data,
	}
}

// 头插法
func (l *DLNode) InsertAtHead(data ...int) {
	for i := 0; i < len(data); i++ {
		n := NewDLNode(data[i])
		if l.IsEmpty() {
			n.Next = l.Next
			n.Prior = l
			l.Next = n
		} else {
			n.Next = l.Next
			n.Prior = l
			l.Next.Prior = n
			l.Next = n
		}
	}
}

// 尾插法
func (l *DLNode) InsertAtTail(data ...int) {
	list := l
	for i := 0; i < len(data); i++ {
		for list.Next != nil {
			list = list.Next
		}
		n := NewDLNode(data[i])
		list.Next = n
		n.Prior = list
	}
}

// 查找
func (l *DLNode) FindAt(n int) (int, error) {
	if n <= 0 {
		return -1, errors.New("n too small")
	}
	if l.IsEmpty() {
		return -1, errors.New("list is empty")
	}
	idx := 0
	list := l
	for list.Next != nil && idx < n {
		list = list.Next
		idx++
	}
	if idx != n {
		return -1, errors.New("n too big")
	}
	return list.Data, nil
}

// 删除
func (l *DLNode) DeleteAt(n int) (error) {
	if n <= 0 {
		return errors.New("n too small")
	}
	idx := 0
	list := l
	var prev *DLNode
	for list.Next != nil && idx < n {
		prev = list
		list = list.Next
		idx++
	}
	if idx != n {
		return errors.New("n too big")
	}
	if list.Next == nil {
		prev.Next = nil
	} else {
		list.Next.Prior = prev
		prev.Next = list.Next
	}
	return nil
}

// 是否为空
func (l *DLNode) IsEmpty() bool {
	return l.Next == nil
}

// 重置
func (l *DLNode) Reset() {
	l.Next = nil
}


func (l *DLNode) Print() []int {
	list := l
	var nums []int
	for list.Next != nil {
		nums = append(nums, list.Next.Data)
		list = list.Next
	}
	return nums
}