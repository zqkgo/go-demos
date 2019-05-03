package dsa

import "testing"

func TestLinkedList(t *testing.T) {
	l := NewLNode(0)
	l.InsertAtHead(3, 23, 55, 78)
	n, _ := l.FindAt(3)
	if n != 23 {
		t.Fatalf("want %d, got %d\n", 23, n)
	}

	l.InsertAtTail(99, 100)
	n, _ = l.FindAt(6)
	if n != 100 {
		t.Fatalf("want %d, got %d\n", 100, n)
	}
	n, _ = l.FindAt(1)
	if n != 78 {
		t.Fatalf("want %d, got %d\n", 78, n)
	}

	_, err := l.FindAt(10)
	if err == nil {
		t.Fatalf("want %s, got %s\n", "n too big", err)
	}

	l.Reset()
	if !l.IsEmpty() {
		t.Fatal("want empty but not")
	}

	l.InsertAtTail(1, 2, 3, 4, 5)
	l.DeleteAt(3)
	n, _ = l.FindAt(3)
	if n != 4 {
		t.Fatalf("want %d, got %d\n", 4, n)
	}
	l.DeleteAt(4) // delete last
	n, _ = l.FindAt(3)
	if n != 4 {
		t.Fatalf("want %d, got %d\n", 4, n)
	}
	l.DeleteAt(1) // delete first
	n, _ = l.FindAt(1)
	if n != 2 {
		t.Fatalf("want %d, got %d\n", 2, n)
	}

	err = l.DeleteAt(3)
	if err == nil {
		t.Fatalf("want %s, got %s\n", "n too big", err)
	}
}
