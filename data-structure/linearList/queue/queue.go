package main

import "fmt"

type Queue struct {
	E        []int
	Front    int
	Rear     int
	Capacity int
}

func main() {
	q := &Queue{
		E:        make([]int, 10),
		Front:    0,
		Rear:     0,
		Capacity: 10,
	}
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	q.EnQueue(7)
	q.EnQueue(8)
	q.EnQueue(9)
	q.EnQueue(10)
	q.EnQueue(11)
	q.EnQueue(12)

	q.PrintQueue()
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())

	q.EnQueue(9)
	q.EnQueue(10)
	q.EnQueue(11)
	q.EnQueue(12)
	q.PrintQueue()
	fmt.Println(q.Front, q.Rear)
}

func (q *Queue) EnQueue(e int) {
	pos := (q.Rear + 1) % q.Capacity
	if pos == q.Front {
		fmt.Println("queue is already full")
		return
	}
	q.Rear = pos
	q.E[q.Rear] = e
}

func (q *Queue) DeQueue() int {
	if q.IsEmpty() {
		return -999
	}
	q.Front = (q.Front + 1) % q.Capacity
	return q.E[q.Front]
}

func (q *Queue) IsEmpty() bool {
	fmt.Println("queue is already empty")
	return q.Front == q.Rear
}

func (q *Queue) PrintQueue() {
	fmt.Print("<<< ")
	i := q.Front
	for i != q.Rear {
		i = (i + 1) % q.Capacity
		fmt.Print(q.E[i], " ")
	}
	fmt.Println("<<<")
}
