package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	make(map[*Node]bool)
	var newHead *Node
	for head != nil {

	}
}
