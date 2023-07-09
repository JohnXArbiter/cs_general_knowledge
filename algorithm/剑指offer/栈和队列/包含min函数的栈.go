package main

type MinStack struct {
	st    []int
	minSt []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.minSt) == 0 {
		this.minSt = append(this.minSt, x)
		this.st = append(this.st, x)
		return
	}
	top := this.minSt[len(this.minSt)-1]
	if x < top {
		this.minSt = append(this.minSt, x)
	} else {
		this.minSt = append(this.minSt, top)
	}
	this.st = append(this.st, x)
}

func (this *MinStack) Pop() {
	this.st = this.st[:len(this.st)-1]
	this.minSt = this.minSt[:len(this.minSt)-1]
}

func (this *MinStack) Top() int {
	return this.st[len(this.st)-1]
}

func (this *MinStack) Min() int {
	return this.minSt[len(this.minSt)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

// 我的拉基写法
//type MinStack struct {
//	Stack []int
//	Minp  int
//	Topp  int
//}
//
///** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{
//		Stack: make([]int, 0),
//		Minp:  -1,
//		Topp:  -1,
//	}
//}
//
//func (this *MinStack) Push(x int) {
//	if this.Topp == -1 {
//		this.Minp = 0
//		this.Topp++
//		this.Stack = append(this.Stack, x)
//	} else {
//		this.Topp++
//		this.Stack = append(this.Stack, x)
//		if x < this.Stack[this.Minp] {
//			this.Minp = this.Topp
//		}
//	}
//}
//
//func (this *MinStack) Pop() {
//	if this.Topp == -1 {
//		return
//	}
//	this.Topp--
//	this.Stack = this.Stack[:len(this.Stack)-1]
//	min := 0
//	for i := 0; i <= this.Topp; i++ {
//		if this.Stack[i] < this.Stack[min] {
//			min = i
//		}
//	}
//	this.Minp = min
//}
//
//func (this *MinStack) Top() int {
//	return this.Stack[this.Topp]
//}
//
//func (this *MinStack) Min() int {
//	return this.Stack[this.Minp]
//}
