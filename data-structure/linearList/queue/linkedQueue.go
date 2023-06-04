package main

import "fmt"

type LQNode struct {
	E    int
	Next *LQNode
}

type LQueue struct {
	Front *LQNode
	Rear  *LQNode
}

func main() {
	lqnode := new(LQNode)
	lq := &LQueue{
		Front: lqnode,
		Rear:  lqnode,
	}
	lq.EnQueue(2)
	lq.EnQueue(453)
	lq.EnQueue(3)
	lq.EnQueue(3)
	lq.EnQueue(3)
	lq.EnQueue(3)
	lq.EnQueue(3)
	lq.EnQueue(3)
	lq.EnQueue(3)
	lq.EnQueue(3)
	lq.EnQueue(3)
	lq.EnQueue(3)
	lq.PrintQueue()
	lq.DeQueue()
	lq.PrintQueue()

}

func (lq *LQueue) EnQueue(e int) {
	lqnode := &LQNode{
		E:    e,
		Next: nil,
	}
	lq.Rear.Next = lqnode
	lq.Rear = lqnode
}

func (lq *LQueue) DeQueue() int {
	if lq.IsEmpty() {
		fmt.Println("linked queue is already empty")
		return -999
	}
	tmp := lq.Front.Next
	lq.Front.Next = lq.Front.Next.Next
	if lq.Rear == tmp {
		lq.Rear = lq.Front
	}
	tmp.Next = nil
	return tmp.E
}

func (lq *LQueue) IsEmpty() bool {
	return lq.Front == lq.Rear
}

func (lq *LQueue) PrintQueue() {
	fmt.Print("<<< ")
	node := lq.Front
	for node != lq.Rear {
		node = node.Next
		fmt.Print(node.E, " ")
	}
	fmt.Println("<<<")
}
