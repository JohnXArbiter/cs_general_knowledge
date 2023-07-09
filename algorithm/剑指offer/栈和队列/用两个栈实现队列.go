package main

func main() {

}

type CQueue struct {
	inStack, outStack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.outStack) == 0 {
		if len(this.inStack) == 0 {
			return -1
		}
		i := len(this.inStack) - 1
		for i >= 0 {
			this.outStack = append(this.outStack, this.inStack[i])
			this.inStack = this.inStack[:len(this.inStack)-1]
			i--
		}
	}
	value := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return value
}
