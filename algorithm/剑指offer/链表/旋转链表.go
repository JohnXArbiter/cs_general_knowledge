package main

import "fmt"

func main() {
	head := &ListNode3{
		Val:  12,
		Next: nil,
	}
	head2 := &ListNode3{
		Val:  265,
		Next: nil,
	}
	head3 := &ListNode3{
		Val:  12,
		Next: nil,
	}
	head4 := &ListNode3{
		Val:  4,
		Next: nil,
	}
	head.Next = head2
	head2.Next = head3
	head3.Next = head4
	rotateRight(head, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

type ListNode3 struct {
	Val  int
	Next *ListNode3
}

// 思路一：每一次循环就是把最后一个节点替换之前的头结点，所以只要替换 k % length次
func rotateRight(head *ListNode3, k int) *ListNode3 {
	if head == nil || k == 0 || head.Next == nil {
		return head
	}
	node := head
	count := 0
	for node != nil {
		count++
		node = node.Next
	}
	newHead, oldHead := head, head
	for i := 0; i < k%count; i++ {
		rear := oldHead
		for rear.Next.Next != nil {
			rear = rear.Next
		}
		newHead = rear.Next
		rear.Next = nil
		newHead.Next = oldHead
		oldHead = newHead
	}
	return newHead
}

// 思路二：将链表先连接成一个循环链表，然后找到移动后的头结点的前一个，断开，返回新的这个头即可
func rotateRight2(head *ListNode3, k int) *ListNode3 {
	if head == nil || k == 0 || head.Next == nil {
		return head
	}
	node := head
	count := 1
	for node.Next != nil {
		count++
		node = node.Next
	}
	if k == count {
		return head
	}
	node.Next = head
	i := count - k%count
	for i > 1 {
		i--
		head = head.Next
	}
	res := head.Next
	head.Next = nil
	return res
}
