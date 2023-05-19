package main

import (
	"fmt"
)

type Node struct {
	data any
	next *Node
}

type LinkedList struct {
	head   *Node
	rear   *Node
	length int
}

func main() {
	fmt.Println("----------------------------------")
	var array = []int{2, 3, 4, 5, 6, 7, 8}
	// 头插法
	fmt.Println("头插法：")
	linkedList1 := initList()
	for _, value := range array {
		hadd(value, linkedList1)
	}
	getAll(linkedList1)
	get(2, linkedList1)
	remove(linkedList1, 3)
	getAll(linkedList1)
	destroy(linkedList1)
	fmt.Println("----------------------------------")

	// 尾插法
	fmt.Println("尾插法:")
	linkedList2 := initList()
	for _, value := range array {
		radd(value, linkedList2)
	}
	getAll(linkedList2)
	get(2, linkedList2)
	//remove(linkedList2, 3)
	insert(3, 2131, linkedList2)
	getAll(linkedList2)
	destroy(linkedList2)
}

func initList() (linkedList *LinkedList) {
	linkedList = new(LinkedList)
	var ptr Node
	ptr.data = nil
	ptr.next = nil
	var l LinkedList
	l.head = &ptr
	l.rear = &ptr
	l.length = 0
	linkedList = &l
	return linkedList
}

func hadd(data int, list *LinkedList) {
	ptr := &Node{
		data: data,
		next: nil,
	}
	ptr.next = list.head.next
	list.head.next = ptr
	list.length++
}

func radd(data int, list *LinkedList) {
	ptr := &Node{
		data: data,
		next: nil,
	}
	list.rear.next = ptr
	list.rear = ptr
	list.length++
}

func get(index int, list *LinkedList) {
	if list.head.next == nil {
		fmt.Println("获取操作：链表没有数据！")
		return
	}
	ptr := list.head.next
	for i := 1; i <= index; i++ {
		if i == index {
			fmt.Printf("获取操作：第%d位的数据是：%v\n", index, ptr.data)
			return
		}
		ptr = ptr.next
	}
}

func getAll(list *LinkedList) {
	data := make([]int, list.length)
	p := list.head.next
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

func remove(list *LinkedList, index int) {
	if index <= 0 || index > list.length {
		fmt.Println("删除操作：索引错误！")
		return
	}
	p := list.head
	if p.next == nil {
		fmt.Println("删除操作：链表没有数据！")
		return
	}
	i := 0
	for p.next != nil {
		i++
		if i == index {
			p.next = p.next.next
			break
		}
		p = p.next
	}
	list.length--
}

func insert(index int, data int, list *LinkedList) {
	if index <= 0 || index > list.length {
		fmt.Println("插入操作：索引错误！")
		return
	}
	p := list.head
	if p.next == nil {
		fmt.Println("插入操作：链表没有数据！")
		return
	}
	for i := 1; i < list.length; i++ {
		p = p.next
	}
	node := &Node{
		data: data,
		next: nil,
	}
	node.next = p.next
	p.next = node
	list.length++
}

func destroy(list *LinkedList) {
	list = nil
}
