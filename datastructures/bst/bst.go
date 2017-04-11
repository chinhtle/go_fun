package main

import "fmt"

type element struct {
	key         int
	left, right *element
}

// BinarySearchTree represents a binary search tree
type BinarySearchTree struct {
	root  *element
	count int
}

func newElement(v int) *element { return &element{v, nil, nil} }

func (b *BinarySearchTree) addHelper(e *element, k int) {
	if b.root == nil {
		b.root = newElement(k)
		b.count++
	} else if k < e.key {
		if e.left != nil {
			b.addHelper(e.left, k)
		} else {
			e.left = newElement(k)
			b.count++
		}
	} else if k > e.key {
		if e.right != nil {
			b.addHelper(e.right, k)
		} else {
			e.right = newElement(k)
			b.count++
		}
	} else {
		fmt.Println("Key", k, "already exists.")
	}
}

func (b *BinarySearchTree) deleteHelper(e *element, k int) *element {
	if e == nil {
		return nil
	} else if k > e.key {
		e.right = b.deleteHelper(e.right, k)
	} else if k < e.key {
		e.left = b.deleteHelper(e.left, k)
	} else {
		deleted := true

		if e.left == nil && e.right == nil {
			e = nil
		} else if e.left == nil {
			r := e.right
			e.right = nil
			e = r
		} else if e.right == nil {
			r := e.left
			e.left = nil
			e = r
		} else {
			deleted = false
			for l := e.right; l != nil; l = l.left {
				if l.left == nil {
					e.key = l.key
				}
			}

			e.right = b.deleteHelper(e.right, e.key)
		}

		if deleted {
			b.count--
		}
	}

	return e
}

func (b *BinarySearchTree) printHelper(e *element) {
	if e == nil {
		return
	}

	b.printHelper(e.left)
	fmt.Print(e.key, " ")
	b.printHelper(e.right)
}

// NewBst returns a new BinarySearchTree
func NewBst() *BinarySearchTree { return &BinarySearchTree{} }

// Len returns the number of elements in the BST
func (b *BinarySearchTree) Len() int { return b.count }

// Add inserts a new element with the key k into the BST
func (b *BinarySearchTree) Add(k int) {
	b.addHelper(b.root, k)
}

// Delete removes an element, with the key k, from the BST
func (b *BinarySearchTree) Delete(k int) {
	if b.root == nil {
		fmt.Println("Empty tree. Cannot delete")
	} else {
		b.root = b.deleteHelper(b.root, k)
	}
}

// Print displays an in-order print of all elements in the BST
func (b *BinarySearchTree) Print() {
	fmt.Println("Tree contents:")
	b.printHelper(b.root)
	fmt.Println()
}

func main() {
	t := NewBst()
	t.Add(2)
	t.Add(1)
	t.Add(4)
	t.Add(3)
	t.Print()
	fmt.Println("Size of tree:", t.Len())

	d := 4
	fmt.Println("Deleting:", d)
	t.Delete(d)
	t.Print()

	d = t.root.key
	fmt.Println("Deleting root:", d)
	t.Delete(d)
	t.Print()

	d = t.root.key
	fmt.Println("Deleting root:", d)
	t.Delete(d)
	t.Print()
}
