package main

import "testing"

func TestBinarySearchTree_Len(t *testing.T) {
	b := NewBst()

	b.count++
	if b.Len() != b.count {
		t.Error("Expecting", b.count, "but received", b.Len())
	}
}

func TestBinarySearchTree_Add(t *testing.T) {
	b := NewBst()
	k := 5

	if b.Add(k); b.root == nil || b.root.key != k {
		t.Error("Invalid root. Expecting", k)
	}

	if b.count != 1 {
		t.Error("Expected count to be 1, received", b.count)
	}

	k = 10
	if b.Add(k); b.root.right.key != k {
		t.Error("Unexpected right element key. Expecting", k)
	}

	k = 1
	if b.Add(k); b.root.left.key != k {
		t.Error("Unexpected left element key. Expecting", k)
	}

	k = 3
	if b.Add(k); b.root.left.right.key != k {
		t.Error("Element with key", k, "not added in the correct location")
	}
}

func TestBinarySearchTree_Delete(t *testing.T) {
	b := NewBst()

	b.Add(2)
	b.Add(1)
	b.Add(4)
	b.Add(3)
	b.Add(5)

	n := b.root.right.right.key
	if b.Delete(4); b.root.right.key != n {
		t.Error("Deleting parent with two leaves failed. Expecting", n)
	}

	if b.count != 4 {
		t.Error("Count after deleting incorrect. Received", b.count)
	}

	n = b.root.right.left.key
	if b.Delete(5); b.root.right.key != n {
		t.Error("Deleting parent with one leaf failed. Expecting", n)
	}

	n = b.root.right.key
	if b.Delete(2); b.root.key != n {
		t.Error("Deleting root with two leafs failed. Expecting", n)
	}

	n = b.root.left.key
	if b.Delete(3); b.root.key != n {
		t.Error("Deleting root with one leaf failed. Expecting", n, "got", b.root.key)
	}

	if b.Delete(1); b.root != nil {
		t.Error("Deleting one element tree failed. Received non-nil root")
	}
}
