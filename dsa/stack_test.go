package dsa

import "testing"

func TestSqStack(t *testing.T) {
	sqStack := NewSqStack(5)
	if !sqStack.IsEmpty() {
		t.Fatal("want empty but not")
	}
	sqStack.Push(1)
	sqStack.Push(2)
	sqStack.Push(3)
	sqStack.Push(4)
	sqStack.Push(5)
	err := sqStack.Push(6)
	if err == nil {
		t.Fatal("want error, but got nil")
	}
	data, _ := sqStack.Pop()
	if data != 5 {
		t.Fatalf("want 5, but got %d\n", data)
	}
	sqStack.Pop()
	sqStack.Pop()
	sqStack.Pop()
	sqStack.Pop()
	_, err = sqStack.Pop()
	if err == nil {
		t.Fatal("want error, but got nil")
	}
}
