package linkedlist

import "errors"

// ErrEmptyList error variable
var ErrEmptyList error = errors.New("Empty LIst")

// NewList function
func NewList(vv ...interface{}) *List {

	l := &List{}
	for _, v := range vv {
		l.PushBack(v)
	}

	return l
}

// Node struct
type Node struct {
	Val  interface{}
	next *Node
	prev *Node
	list *List
}

// Next function
func (n Node) Next() *Node {

	if n.list.reversed {
		return n.prev
	}
	return n.next
}

// Prev function
func (n Node) Prev() *Node {

	if n.list.reversed {
		return n.next
	}
	return n.prev
}

// First function
func (n Node) First() *Node {

	return n.list.First()
}

// Last function
func (n Node) Last() *Node {

	return n.list.Last()
}

// List struct
type List struct {
	first    *Node
	last     *Node
	reversed bool
}

// First function
func (l List) First() *Node {
	if l.reversed {
		return l.last
	}
	return l.first
}

// Last function
func (l List) Last() *Node {
	if l.reversed {
		return l.first
	}
	return l.last
}

// PushBack function
func (l *List) PushBack(v interface{}) {

	n := &Node{Val: v, list: l}

	if l.first == nil {
		l.first = n
		l.last = n
		return
	}

	// Set the old tail to the new node
	n.prev = l.last
	// Point old tail 'next' to the new node
	l.last.next = n
	// Set the new tail
	l.last = n
}

// PopBack function
func (l *List) PopBack() (n interface{}, err error) {

	if l.last == nil {
		return 0, ErrEmptyList
	}

	n = l.last.Val

	l.last = l.last.prev
	if l.last == nil {
		l.first = nil
		return
	}
	l.last.next = nil

	return
}

// PushFront function
func (l *List) PushFront(v interface{}) {

	n := &Node{Val: v, list: l}

	if l.first == nil {
		l.first = n
		l.last = n
		return
	}

	// Set the old head to the new node
	n.next = l.first
	// Point old head 'prev' to the new node
	l.first.prev = n
	// Set the new head
	l.first = n
}

// PopFront function
func (l *List) PopFront() (n interface{}, err error) {

	if l.first == nil {
		return 0, ErrEmptyList
	}

	n = l.first.Val

	l.first = l.first.next
	if l.first == nil {
		l.last = nil
		return
	}
	l.first.prev = nil

	return
}

// Reverse function
func (l *List) Reverse() {

	l.reversed = !l.reversed
}
