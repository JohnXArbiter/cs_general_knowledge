package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	f, r := head, head.Next
	for r != nil {
		if r.Val == val {
			f.Next = r.Next
			break
		}
		f, r = f.Next, r.Next
	}
	return head
}
