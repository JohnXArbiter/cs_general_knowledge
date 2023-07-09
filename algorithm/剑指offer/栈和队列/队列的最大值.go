package main

type MaxQueue struct {
	Data []int
	MaxQ []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.MaxQ) == 0 {
		return -1
	}
	return this.MaxQ[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.Data = append(this.Data, value)
	if len(this.MaxQ) != 0 && value > this.MaxQ[len(this.MaxQ)-1] {
		var l = len(this.MaxQ) - 1
		for ; l >= 0; l-- {
			if value < this.MaxQ[l] {
				break
			}
		}
		this.MaxQ = this.MaxQ[:l+1]
	}
	this.MaxQ = append(this.MaxQ, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.Data) == 0 {
		return -1
	}
	res := this.Data[0]
	this.Data = this.Data[1:]
	if res == this.MaxQ[0] {
		this.MaxQ = this.MaxQ[1:]
	}
	return res
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
