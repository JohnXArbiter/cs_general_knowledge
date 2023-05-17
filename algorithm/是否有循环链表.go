package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var map111 = map[*ListNode]bool{}

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}

	l2 := &ListNode{
		Val:  1,
		Next: l1,
	}

	l3 := &ListNode{
		Val:  1,
		Next: l2,
	}

	l4 := &ListNode{
		Val:  1,
		Next: l3,
	}
	fmt.Printf("%v\n", l4)
	l1.Next = l4
	fmt.Printf("%t", hasCycle(l1))
}

func hasCycle(head *ListNode) bool {
	var p = new(ListNode)
	p.Next = new(ListNode)
	p = head
	if p.Next == nil {
		return false
	}
	for {
		if map111[p] {
			fmt.Printf("打印p:%v\n", p)
			fmt.Printf("打印*p:%v\n", *p)

			return true
		}
		map111[p] = true
		if p.Next == nil {
			return false
		}
		p = p.Next
	}
}
