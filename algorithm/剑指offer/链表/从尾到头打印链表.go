package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var a []int
	for head != nil {
		a = append(a, head.Val)
		head = head.Next
	}
	for i, j := 0, len(a)-1; i < j; {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	return a
}
