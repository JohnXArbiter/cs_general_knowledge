package main

import "fmt"

func main() {
	//af := new(abstractFactory)

}

// 三个功能
type abstractGPU interface {
	display()
}

type abstractMemory interface {
	storage()
}

type abstractCPU interface {
	calculate()
}

// 三个厂商的三种产品
// intel的
type intelCPU struct{}

func (*intelCPU) calculate() {
	fmt.Println("intel's cpu is calculating...")
}

type intelMemory struct{}

func (*intelMemory) storage() {
	fmt.Println("intel's memory is storing...")
}

type intelGPU struct{}

func (*intelGPU) display() {
	fmt.Println("intel's gpu is displaying...")
}

type kingstonCPU struct{}

func (*kingstonCPU) calculate() {
	fmt.Println("kingston's cpu is calculating...")
}

type kingstonMemory struct{}

func (*kingstonMemory) storage() {
	fmt.Println("kingston's memory is storing...")
}

type kingstonGPU struct{}

func (*kingstonGPU) display() {
	fmt.Println("kingston's gpu is displaying...")
}

type nvidiaCPU struct{}

func (*nvidiaCPU) calculate() {
	fmt.Println("nvidia's cpu is calculating...")
}

type nvidiaMemory struct{}

func (*nvidiaMemory) storage() {
	fmt.Println("nvidia's memory is storing...")
}

type nvidiaGPU struct{}

func (*nvidiaGPU) display() {
	fmt.Println("nvidia's gpu is displaying...")
}

// 抽象工厂
type abstractFactory interface {
	createCPU()
	createMemory()
	createGPU()
}

// Intel厂商
type intelFactory struct {
}

func (*intelFactory) createCPU() abstractCPU {
	var cpu abstractCPU
	cpu = new(intelCPU)
	return cpu
}

func (*intelFactory) createMemory() {

}
func (*intelFactory) createGPU() {

}

// nvidia厂商
type nvidiaFactory struct {
}

func (*nvidiaFactory) createCPU() {

}

func (*nvidiaFactory) createMemory() {

}
func (*nvidiaFactory) createGPU() {

}

// kingston厂商
type kingstonFactory struct {
}

func (*kingstonFactory) createCPU() {

}

func (*kingstonFactory) createMemory() {

}
func (*kingstonFactory) createGPU() {

}
