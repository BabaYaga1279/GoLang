package lru

import (
	"fmt"
)

type LRU_Queue struct {
	queue Queue
	hash  map[int]*Node
	cap   int
}

// create new LRU_Queue object
func NewLRU_Queue(cap int) LRU_Queue {
	return LRU_Queue{
		NewQueue(),
		make(map[int]*Node),
		cap,
	}
}

func Test() {
	fmt.Println("LRU")
}

// manualy set frame cap, maximum number of process in the frame before lru get replaced
func (q *LRU_Queue) SetCap(cap int) {
	q.cap = cap
}

func (q LRU_Queue) GetLen() int {
	return q.queue.len
}

func (q LRU_Queue) CapReached() bool {
	if q.GetLen() == q.cap {
		return true
	}

	return false
}

func (q LRU_Queue) Top() int {
	if q.GetLen() == 0 {
		return -1
	}
	return q.queue.first.val
}

func (q *LRU_Queue) pop() int {
	top := q.Top()
	if top < 0 {
		return top
	}

	q.queue.Pop()
	q.hash[top] = nil

	return top
}

func (q *LRU_Queue) push(val int) {
	node := NewNode(val)

	q.queue.PushNode(node)
	q.hash[val] = node
}

// when a process is used, move it to back of the queue
func (q *LRU_Queue) moveToBack(val int) {
	q.hash[val] = q.queue.MoveNodeToBack(q.hash[val])
}

// use a process
func (q *LRU_Queue) Use(val int) bool {
	node := q.hash[val]
	ok := bool(node != nil)

	//fmt.Println("Hit: ", ok)

	if ok == false {
		if q.CapReached() {
			q.pop()
			//fmt.Println("poped: ", t, q.hash[t])
		}

		q.push(val)
	} else {
		q.moveToBack(val)
	}

	return ok
}

func (q LRU_Queue) PrintLRU_Queue() {
	q.queue.PrintQueue()
}
