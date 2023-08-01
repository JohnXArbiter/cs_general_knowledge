package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 != nil && l2 == nil {
		return l1
	} else if l2 != nil && l1 == nil {
		return l2
	}
	var head, l *ListNode
	if l1.Val <= l2.Val {
		head, l = l1, l1
		l1 = l1.Next
	} else {
		head, l = l2, l2
		l2 = l2.Next
	}
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			l.Next = l1
			l1 = l1.Next
		} else {
			l.Next = l2
			l2 = l2.Next
		}
		l = l.Next
	}
	if l1 != nil {
		l.Next = l1
	} else {
		l.Next = l2
	}
	return head
}

/**
 * 大佬的解
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil { //递归终止条件
		return l2
	}
	if l2 == nil { //递归终止条件
		return l1
	}
	if l1.Val > l2.Val { //交换两节点
		l1, l2 = l2, l1
	}
	l1.Next = mergeTwoLists(l1.Next, l2) //递归
	return l1
}
