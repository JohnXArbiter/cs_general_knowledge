package main

import "fmt"

var instance = new(singleton)

type singleton struct {
}

func GetInstance() *singleton {
	return instance
}

func (s *singleton) doSomething() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstance()
	s.doSomething()
}
