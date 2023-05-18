package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

var instancecf *singletoncf

type singletoncf struct {
}

func GetInstancecf() *singletoncf {
	lock.Lock()
	defer lock.Unlock()
	if instancecf == nil {
		instancecf = new(singletoncf)
	}
	return instancecf
}

func (s *singletoncf) doSomething() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstancecf()
	s.doSomething()
}
