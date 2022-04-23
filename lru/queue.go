package lru

import (
	"fmt"
)

type Node struct {
	val        int
	prev, next *Node
}

type Queue struct {
	first, last *Node
	len         int
}

func NewNode(val int) *Node {
	return &Node{val, nil, nil}
}

func NewQueue() Queue {
	return Queue{nil, nil, 0}
}

func (q *Queue) PushNode(node *Node) {
	if node == nil {
		return
	}

	q.len++

	node.next = nil
	node.prev = nil

	if q.first == nil {
		q.first = node
		q.last = node
		return
	}

	q.last.next = node
	node.prev = q.last
	q.last = node

	return
}

func (q *Queue) PushVal(val int) {
	node := NewNode(val)

	q.PushNode(node)

	return
}

func (q *Queue) Pop() *Node {
	if q.first == nil {
		return nil
	}

	q.len--

	node := q.first

	if q.first == q.last {
		q.first = nil
		q.last = nil
		return node
	}

	q.first = q.first.next
	q.first.prev.next = nil
	q.first.prev = nil

	return node
}

func (q *Queue) PopLast() *Node {
	if q.first == nil {
		return nil
	}

	node := q.last

	if q.first == q.last {
		return q.Pop()
	}

	q.len--

	q.last = q.last.prev
	q.last.next.prev = nil
	q.last.next = nil

	return node
}

func (q *Queue) popNode(node *Node) *Node {
	if node == nil {
		return nil
	}

	if q.first == q.last {
		return q.Pop()
	}

	if node == q.first {
		return q.Pop()
	}

	if node == q.last {
		return q.PopLast()
	}

	q.len--

	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
	node.next = nil
	node.prev = nil

	return node
}

func (q Queue) Top() *Node {
	return q.first
}

func (q Queue) Bottom() *Node {
	return q.last
}

func (q Queue) PrintQueue() {
	for i := q.first; i != nil; i = i.next {
		fmt.Print(i.val, " ")
	}

	fmt.Println("")
}

// move a node to its back feature
func (q *Queue) MoveNodeToBack(node *Node) *Node {
	if node == nil || q.first == nil || q.first == q.last {
		return q.first
	}

	node = q.popNode(node)
	q.PushNode(node)

	return q.last
}
