package main

/*
双队列实现：
一个队列用来入栈，另一个队列保存之前入栈的元素
每次都将新入栈的元素放入暂存队列的头部，这样能保证新入栈的元素永远在前面
*/
type MyStack struct {
	Queue1 []int
	Queue2 []int
}

func Constructor2() MyStack {
	return MyStack{
		Queue1: make([]int, 0),
		Queue2: make([]int, 0),
	}
}

func (s *MyStack) Push(x int) {
	s.Queue1 = append(s.Queue1, x)
	s.move()
}

func (s *MyStack) move() {
	if len(s.Queue2) == 0 {
		s.Queue1, s.Queue2 = s.Queue2, s.Queue1
	} else {
		s.Queue1 = append(s.Queue1, s.Queue2[0])
		s.Queue2 = s.Queue2[1:]
		s.move()
	}
}

func (s *MyStack) Pop() int {
	res := s.Queue2[0]
	s.Queue2 = s.Queue2[1:]
	return res
}

func (s *MyStack) Top() int {
	return s.Queue2[0]
}

func (s *MyStack) Empty() bool {
	return len(s.Queue2) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

/*
单队列，不知道算不算作弊
*/
type MyStack struct {
	Queue []int
}

func Constructor() MyStack {
	return MyStack{Queue: make([]int, 0)}
}

func (this *MyStack) Push(x int) {
	tmp := make([]int, 0)
	tmp = append(tmp, x)
	tmp = append(tmp, this.Queue...)
	this.Queue = tmp
}

func (this *MyStack) Pop() int {
	res := this.Queue[0]
	this.Queue = this.Queue[1:]
	return res
}

func (this *MyStack) Top() int {
	return this.Queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.Queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
