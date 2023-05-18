package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 标记
var initialized uint32
var locka sync.Mutex

type singeltona struct{}

var instancea *singeltona

func GetInstancea() *singeltona {
	//如果标记为被设置，直接返回，不加锁
	if atomic.LoadUint32(&initialized) == 1 {
		return instancea
	}

	//如果没有，则加锁申请
	locka.Lock()
	defer locka.Unlock()

	if initialized == 0 {
		instancea = new(singeltona)
		//设置标记位
		atomic.StoreUint32(&initialized, 1)
	}

	return instancea
}

func (s *singeltona) doSomething() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstancea()
	s.doSomething()
}
