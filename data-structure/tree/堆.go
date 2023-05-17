package main

import "fmt"

type MaxHeap struct {
	arr      []int
	size     int
	capacity int
}

func main() {
	heap := &MaxHeap{
		arr:      make([]int, 11),
		size:     0,
		capacity: 10,
	}
	heap.size = 0
	InsertH(heap, 3)
	InsertH(heap, 456)
	InsertH(heap, 32)
	InsertH(heap, 345)
	InsertH(heap, 53)
	InsertH(heap, 31)
	fmt.Println(heap, heap.arr)
}

func InsertH(heap *MaxHeap, e int) {
	// 满了退出
	if heap.size == heap.capacity {
		return
	}
	index := heap.size
	index++ // 获取插入的位置，肯定是之前的size➕1
	// 条件1是长度大于1的时候才执行，因为size为0和1的时候本来就是直接插入
	// 调减2是当前要插入的值是不是大于他要插入的位置的父亲节点的值，是就交换位置
	for index > 1 && e > heap.arr[index/2] {
		heap.arr[index] = heap.arr[index/2] // 此处的交换位置，是类似插入排序那样，先将之前的值后移（下移）
		index /= 2                          // 不断除2去找更上次的节点
	}
	heap.arr[index] = e
}

func DeleteH(heap *MaxHeap) int {
	if heap.size == 0 {
		return -1
	}
	max, e := heap.arr[1], heap.arr[heap.size-1] // 直接出掉第一个元素就是最大值，然后保存最后一个值
	heap.size--                                  // 出掉了一个所以size--
	index := 1
	// 开始从根节点搜索子树的根节点的大小，乘2才能得到，所以肯定不能大于size
	for index*2 <= heap.size {
		child := index * 2
		// 判断是左子树根和右子树根哪个大
		if child < heap.size && heap.arr[child] < heap.arr[child+1] {
			child++
		}
		// 如果最后最后的节点都比当前子树根大了，那就说明到了
		if e >= heap.arr[child] {
			break
		} else {
			// 没找到就继续下移
			heap.arr[index] = heap.arr[child]
		}
		index = child
	}
	heap.arr[index] = e
	return max
}
