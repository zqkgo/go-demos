package dsa

import (
	"testing"
)

func TestSqQueue(t *testing.T) {
	queue := NewSqQueue(5)
	if !queue.IsEmpty() {
		t.Fatal("want empty but not")
	}
	queue.EnQueue(10)
	queue.EnQueue(20)
	data, _ := queue.DeQueue()
	if data != 10 {
		t.Fatalf("want 10, got %d\n", data)
	}
	queue.DeQueue()
	if !queue.IsEmpty() {
		t.Fatal("want empty but not")
	}

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	if !queue.IsFull() {
		t.Fatal("want full but not")
	}
	err := queue.EnQueue(5)
	if err == nil {
		t.Fatal("want error, got nil")
	}
	data, _ = queue.DeQueue()
	if data != 1 {
		t.Fatalf("want 1, got %d\n", data)
	}
}