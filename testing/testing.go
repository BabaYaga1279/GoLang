package testing

import (
	"fmt"

	"github.com/BabaYaga1279/GoLang/lru"
)

func TestQueue() {
	queue := lru.NewQueue()

	var node *lru.Node = nil

	for i := 1; i <= 5; i++ {
		queue.PushVal(i)
	}

	node = queue.Top()

	fmt.Println("queue: ")
	queue.PrintQueue()

	queue.MoveNodeToBack(node)
	fmt.Println("move 3 to back ")
	queue.PrintQueue()

	queue.Pop()
	fmt.Println("Pop ")
	queue.PrintQueue()

	queue.PopLast()
	fmt.Println("Pop Last ")
	queue.PrintQueue()
}

func TestLRU_Queue() {
	queue := lru.NewLRU_Queue(4)

	input := []int{0, 1, 7, 2, 3, 2, 7, 1, 0, 3}

	for i := 0; i < len(input); i++ {
		fmt.Println("Use: ", input[i])
		queue.Use(input[i])
		queue.PrintLRU_Queue()
	}
}
