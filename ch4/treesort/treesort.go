// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

//!+
type Tree struct {
	value       int
	left, right *Tree
}

func (t *Tree) Left() *Tree  { return t.left }
func (t *Tree) Right() *Tree { return t.right }
func (t *Tree) Value() int   { return t.value }

// Sort sorts values in place.
func Sort(values []int) []int {
	var root *Tree
	for _, v := range values {
		root = add(root, v)
	}
	return appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *Tree, value int) *Tree {
	if t == nil {
		return &Tree{value: value}
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

//!-
