package dsa

import "testing"

func TestBinaryTree(t *testing.T) {
	n1 := NewBTNode(10)
	n2 := NewBTNode(15)
	n3 := NewBTNode(20)
	n4 := NewBTNode(25)
	n5 := NewBTNode(30)
	n6 := NewBTNode(35)
	n7 := NewBTNode(40)
	n1.AddLeft(n2)
	n1.AddRight(n3)
	n2.AddLeft(n4)
	n2.AddRight(n5)
	n3.AddLeft(n6)
	n3.AddRight(n7)

	resetVisits()
	n1.PreOrder()
	visits := visitsStr()
	if visits != "10,15,25,30,20,35,40" {
		t.Fatalf("want '10,15,25,30,20,35,40', got %s\n", visits)
	}

	resetVisits()
	n1.InOrder()
	visits = visitsStr()
	if visits != "25,15,30,10,35,20,40" {
		t.Fatalf("want '25,15,30,10,35,20,40', got %s\n", visits)
	}

	resetVisits()
	n1.PostOrder()
	visits = visitsStr()
	if visits != "25,30,15,35,40,20,10" {
		t.Fatalf("want '25,30,15,35,40,20,10', got %s\n", visits)
	}

	resetVisits()
	n1.Level(8)
	visits = visitsStr()
	if visits != "10,15,20,25,30,35,40" {
		t.Fatalf("want '10,15,20,25,30,35,40', got %s\n", visits)
	}
}
