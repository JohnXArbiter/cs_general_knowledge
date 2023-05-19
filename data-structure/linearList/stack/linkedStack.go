package main

import "fmt"

type LStack struct {
	E    int
	Next *LStack
}

func main() {
	head := new(LStack)
	head.Push(2)
	head.Push(4)

	head.PrintStack()
}

func (ls *LStack) Push(e int) {
	node := &LStack{
		E:    e,
		Next: ls.Next,
	}
	ls.Next = node
}

func (ls *LStack) Pop(head *LStack) int {
	if head.Next == nil {
		fmt.Println("linked stack is already empty")
		return 0
	}
	tmpNode := head.Next
	head.Next = head.Next.Next
	return tmpNode.E
}

func (ls *LStack) IsEmpty(head *LStack) bool {
	if head.Next == nil {
		return true
	}
	return false
}

func (ls *LStack) PrintStack() {
	node := ls.Next
	for node != nil {
		fmt.Print(ls.E, " ")
		node = node.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
