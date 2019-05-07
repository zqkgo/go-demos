package dsa

import (
	"testing"
	"strconv"
	"math/rand"
)

func TestSort(t *testing.T)  {
	sorted := []int{1, 2, 2, 5, 6, 7, 32, 45, 65, 99, 111, 333, 5556}
	elements := []int {1,5,6,7,111,2,333,99,45,65,32,5556,2}
	InsertionSort(elements)
	t1 := joinElements(sorted)
	t2 := joinElements(elements)
	if t1 != t2 {
		t.Fatalf("want %s, got %s\n", t1, t2)
	}

	elements = shuffle(elements)
	BubbleSort(elements)
	t1 = joinElements(sorted)
	t2 = joinElements(elements)
	if t1 != t2 {
		t.Fatalf("want %s, got %s\n", t1, t2)
	}

	elements = shuffle(elements)
	elements = QuickSort(elements)
	t1 = joinElements(sorted)
	t2 = joinElements(elements)
	if t1 != t2 {
		t.Fatalf("want %s, got %s\n", t1, t2)
	}

	elements = shuffle(elements)
	SelectionSort(elements)
	t1 = joinElements(sorted)
	t2 = joinElements(elements)
	if t1 != t2 {
		t.Fatalf("want %s, got %s\n", t1, t2)
	}

	elements = shuffle(elements)
	elements = MergeSort(elements)
	t1 = joinElements(sorted)
	t2 = joinElements(elements)
	if t1 != t2 {
		t.Fatalf("want %s, got %s\n", t1, t2)
	}

	a := []int{1,5,65}
	b := []int{2,111,5556}
	c := merge(a, b)
	if joinElements(c) != "1,2,5,65,111,5556" {
		t.Fatalf("want %s, got %s\n", "1,2,5,65,111,5556", joinElements(c))
	}
}

func joinElements(elements []int) string {
	str := ""
	for _, v := range elements {
		str += strconv.Itoa(v) + ","
	}
	return str[:len(str) - 1]
}

func shuffle(elements []int) []int {
	len := len(elements)
	var newEles []int
	nums := rand.Perm(len)
	for _, n := range nums {
		newEles = append(newEles, elements[n])
	}
	return newEles
}