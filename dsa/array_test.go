package dsa

import "testing"

func TestArray(t *testing.T) {
	arr := NewArray(10)
	arr.InsertAt(1, 10)
	arr.InsertAt(1, 20)
	arr.InsertAt(1, 30)
	arr.InsertAt(3, 40)
	if v, _ := arr.FindAt(4); v != 10 {
		t.Errorf("should be %d but got %d\n", 10, v)
	}

	arr.DeleteAt(2)
	if v, _ := arr.FindAt(2); v != 40 {
		t.Errorf("should be %d but got %d\n", 40, v)
	}

	old, _ := arr.ReplaceAt(2, 66)
	if old != 40 {
		t.Errorf("should be %d but got %d\n", 40, old)
	}
	if v, _ := arr.FindAt(2); v != 66 {
		t.Errorf("should be %d but got %d\n", 66, v)
	}
}
