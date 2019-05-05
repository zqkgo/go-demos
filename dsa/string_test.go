package dsa

import "testing"

func TestString(t *testing.T) {
	p,_ := simplePattern("hello world", "world")
	if p != 6 {
		t.Fatalf("want 6, got %d\n", p)
	}
	_, err := simplePattern("hello world", "world hello universe")
	if err == nil {
		t.Fatal("want error, got nil")
	}

	main := "ABCDABDABCDEFABD"
	sub := "ABDAB"
	p,_ = kmp(main, sub)
	if p != 4 {
		t.Fatalf("want 4, got %d\n", p)
	}
	t.Log(next(sub))
}
