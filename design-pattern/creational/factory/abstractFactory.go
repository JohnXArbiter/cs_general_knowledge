package main

import "fmt"

func main() {
	intelf := new(intelFactory)
	intelc := intelf.createCPU()
	intelc.calculate()

	nvidiaf := new(nvidiaFactory)
	nvidiag := nvidiaf.createGPU()
	nvidiag.display()

	kinsgtonf := new(kingstonFactory)
	kingstonm := kinsgtonf.createMemory()
	kingstonm.storage()
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

// kingston的
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

// nvidia的
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

func (*intelFactory) createMemory() abstractMemory {
	var memory abstractMemory
	memory = new(intelMemory)
	return memory
}

func (*intelFactory) createGPU() abstractGPU {
	var gpu abstractGPU
	gpu = new(intelGPU)
	return gpu
}

// nvidia厂商
type nvidiaFactory struct {
}

func (*nvidiaFactory) createCPU() abstractCPU {
	var cpu abstractCPU
	cpu = new(nvidiaCPU)
	return cpu
}

func (*nvidiaFactory) createMemory() abstractMemory {
	var memory abstractMemory
	memory = new(nvidiaMemory)
	return memory
}
func (*nvidiaFactory) createGPU() abstractGPU {
	var gpu abstractGPU
	gpu = new(nvidiaGPU)
	return gpu
}

// kingston厂商
type kingstonFactory struct {
}

func (*kingstonFactory) createCPU() abstractCPU {
	var cpu abstractCPU
	cpu = new(kingstonCPU)
	return cpu
}

func (*kingstonFactory) createMemory() abstractMemory {
	var memory abstractMemory
	memory = new(kingstonMemory)
	return memory
}
func (*kingstonFactory) createGPU() abstractGPU {
	var gpu abstractGPU
	gpu = new(kingstonGPU)
	return gpu
}
