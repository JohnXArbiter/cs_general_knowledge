package main

import (
	"fmt"
	"sync"
)

var once sync.Once

var instanceOnce *SingletonOnce

type SingletonOnce struct {
}

func GetInstanceOnce() *SingletonOnce {
	once.Do(func() {
		instanceOnce = new(SingletonOnce)
	})
	return instanceOnce
}

func (s *SingletonOnce) doSomething() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstanceOnce()
	s.doSomething()
}
