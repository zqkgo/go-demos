package dsa

import (
	"testing"
	"math/rand"
	"time"
)

func TestSearch(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	idx := BinarySearch(elements, 2)
	if idx != 1 {
		t.Fatalf("want %d, got %d\n", 1, idx)
	}

	idx = Search(elements, 7)
	if idx != 6 {
		t.Fatalf("want %d, got %d\n", 6, idx)
	}

	elements = rand.Perm(1000000)
	elements = QuickSort(elements)
	s1 := time.Now()
	idx1 := BinarySearch(elements, 700034)
	p1 := time.Since(s1)
	s2 := time.Now()
	idx2 := Search(elements, 700034)
	p2 := time.Since(s2)
	if idx1 != idx2 {
		t.Fatalf("want equal but not, idx1=%d, idx2=%d\n", idx1, idx2)
	}
	t.Logf("binary search used: %s, normal search used: %s\n", p1, p2)
}
