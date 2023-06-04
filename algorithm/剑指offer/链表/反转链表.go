package main

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func main() {
	head := &ListNode2{
		Val:  12,
		Next: nil,
	}
	head2 := &ListNode2{
		Val:  265,
		Next: nil,
	}
	head3 := &ListNode2{
		Val:  12,
		Next: nil,
	}
	head4 := &ListNode2{
		Val:  4,
		Next: nil,
	}
	head.Next = head2
	head2.Next = head3
	head3.Next = head4
}

func reverseList(head *ListNode2) *ListNode2 {
	if head == nil {
		return nil
	}
	var tmp, newHead *ListNode2
	for head != nil {
		tmp = head.Next
		head.Next = newHead
		newHead = head
		head = tmp
	}
	return newHead
}
