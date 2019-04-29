package dsa

import (
	"github.com/pkg/errors"
	"fmt"
)

type Array struct {
	Arr   []interface{}
	Used  int
	Total int
}

func NewArray(total int) *Array {
	arr := make([]interface{}, total)
	return &Array{
		Arr:   arr,
		Used:  0,
		Total: total,
	}
}

func (a *Array) InsertAt(n int, val interface{}) error {
	if a.IsFull() {
		return errors.New("array is full")
	}
	if !a.IsNValid(n) {
		return errors.New("invalid n")
	}

	for i := a.Used; i > n-1; i-- {
		a.Arr[i] = a.Arr[i-1]
	}
	a.Arr[n-1] = val

	a.Used++

	return nil
}

func (a *Array) DeleteAt(n int) {
	if a.IsEmpty() {
		return
	}
	if !a.IsNValid(n) {
		return
	}
	if n == a.Used {
		a.Arr[n-1] = nil
	} else {
		for i := n - 1; i < a.Used-1; i++ {
			a.Arr[i] = a.Arr[i+1]
		}
	}

	a.Used--
}

func (a *Array) FindAt(n int) (interface{}, error) {
	if !a.IsNValid(n) {
		return nil, errors.New("invalid n")
	}
	return a.Arr[n-1], nil
}

func (a *Array) ReplaceAt(n int, val interface{}) (interface{}, error) {
	if !a.IsNValid(n) {
		return nil, errors.New("invalid n")
	}
	old := a.Arr[n-1]
	a.Arr[n-1] = val
	return old, nil
}

func (a *Array) IsFull() bool {
	return a.Used == a.Total
}

func (a *Array) IsEmpty() bool {
	return a.Used == 0
}

func (a *Array) Print() {
	for i := 0; i < a.Used; i++ {
		fmt.Println(a.Arr[i])
	}
}

// 限制n范围
func (a *Array) IsNValid(n int) bool {
	if a.Used == 0 {
		return n == 1
	}
	return n > 0 && n <= a.Used
}
