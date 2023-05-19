package main

import "fmt"

type Stack struct {
	E        []int
	Capacity int
	Top      int
}

func main() {
	stack := &Stack{
		E:        make([]int, 0),
		Capacity: 10,
		Top:      -1,
	}
	stack.Push(9)
	stack.Push(23)
	stack.Push(1)
	stack.PrintStack()
	stack.Pop()
	stack.PrintStack()
	stack.Pop()
	stack.PrintStack()
	stack.Pop()
	stack.PrintStack()
	stack.Pop()
	stack.Push(1)
	stack.Push(1)
	stack.Push(1)
	stack.Push(1)
	stack.Push(1)
	stack.Push(1)
	stack.Push(1)
	stack.Push(1)
	stack.Push(1)
	stack.Push(1)
	stack.Push(1)
	stack.Push(1)
	stack.Push(1)
	stack.Push(1)
	stack.Push(1)
	stack.PrintStack()
}

func (s *Stack) Push(e int) {
	if s.Top >= s.Capacity-1 {
		fmt.Println("stack is full")
		return
	}
	s.Top++
	s.E = append(s.E, e)
}

func (s *Stack) Pop() int {
	if s.Top <= -1 {
		fmt.Println("stack is already empty")
		return -999
	}
	e := s.E[s.Top]
	s.Top--
	return e
}

func (s *Stack) PrintStack() {
	if s.Top == -1 {
		fmt.Println("stack is already empty")
		return
	}
	for i := 0; i <= s.Top; i++ {
		fmt.Print(s.E[i], " ")
	}
	fmt.Println()
}

func (s *Stack) IsEmpty() bool {
	if s.Top == -1 {
		fmt.Println("stack is empty")
		return true
	} else {
		fmt.Println("stack is not empty")
		return false
	}
}
