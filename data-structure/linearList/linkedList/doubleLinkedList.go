package main

import "fmt"

type DNode struct {
	data  any
	prior *DNode
	next  *DNode
}

type DLinkedList struct {
	head   *DNode
	rear   *DNode
	length int
}

func main() {
	fmt.Println("----------------------------------")
	var array = []int{2, 3, 4, 5, 6, 7, 8}

	// 尾插法
	fmt.Println("尾插法:")
	dLinkedList := dinitList()
	for _, value := range array {
		dradd(dLinkedList, value)
	}
	dgetAll(dLinkedList)
	dget(dLinkedList, 2)
	dremove(dLinkedList, 3)
	dgetAll(dLinkedList)
	dinsert(dLinkedList, 1, 213)
	dgetAll(dLinkedList)
	dinsert(dLinkedList, 1, 22)
	dgetAll(dLinkedList)
	dinsert(dLinkedList, 1, 11)
	dgetAll(dLinkedList)
	dinsert(dLinkedList, 1, 33)
	dgetAll(dLinkedList)
	ddestroy(dLinkedList)
}

func dinitList() (linkedList *DLinkedList) {
	linkedList = new(DLinkedList)
	var ptr DNode
	ptr.data = nil
	ptr.next = nil
	var l DLinkedList
	l.head = &ptr
	l.rear = &ptr
	l.length = 0
	linkedList = &l
	return linkedList
}

func dradd(dLinkedList *DLinkedList, data any) {
	ptr := &DNode{
		data:  data,
		prior: nil,
		next:  nil,
	}
	ptr.data = data
	if dLinkedList.length == 0 {
		dLinkedList.head.next = ptr
		dLinkedList.rear = ptr
		dLinkedList.rear.prior = dLinkedList.head
	} else {
		ptr.prior = dLinkedList.rear
		dLinkedList.rear.next = ptr
		dLinkedList.rear = ptr
	}
	dLinkedList.length++
}

func dget(dLinkedList *DLinkedList, index int) {
	if dLinkedList.head.next == nil {
		fmt.Println("获取操作：链表没有数据！")
		return
	}
	ptr := dLinkedList.head.next
	for i := 1; i <= index; i++ {
		if i == index {
			fmt.Printf("获取操作：第%d位的数据是：%v\n", index, ptr.data)
			return
		}
		ptr = ptr.next
	}
}

func dgetAll(dLinkedList *DLinkedList) {
	data := make([]int, dLinkedList.length)
	p := dLinkedList.head.next
	if p == nil {
		fmt.Println("全获取操作：链表没有数据！")
		return
	}
	i := 1
	for p != nil {
		data[i-1] = p.data.(int)
		p = p.next
		i++
	}
	fmt.Printf("遍历结果：")
	for _, value := range data {
		fmt.Printf("%d  ", value)
	}
	fmt.Println()
}

func dremove(dLinkedList *DLinkedList, index int) {
	if index <= 0 || index > dLinkedList.length {
		fmt.Println("删除操作：索引错误！")
		return
	}
	p := dLinkedList.head
	if p.next == nil {
		fmt.Println("删除操作：链表没有数据！")
		return
	}
	i := 0
	for p.next != nil {
		i++
		if i == index {
			p.next.next.prior = p
			p.next = p.next.next
			break
		}
		p = p.next
	}
	dLinkedList.length--
}

func dinsert(dLinkedList *DLinkedList, index int, data any) {
	if index <= 0 || index > dLinkedList.length {
		fmt.Println("插入操作：索引错误！")
		return
	}
	p := dLinkedList.head
	if p.next == nil {
		fmt.Println("插入操作：链表没有数据！")
		return
	}
	for i := 1; i < index; i++ {
		p = p.next
	}
	node := &DNode{
		data: data,
		next: nil,
	}
	node.next = p.next
	node.prior = p
	p.next.prior = node
	p.next = node
	dLinkedList.length++
}

func ddestroy(list *DLinkedList) {
	list = nil
}
