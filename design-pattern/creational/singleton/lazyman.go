package main

import "fmt"

var instancel *singletonl

type singletonl struct {
}

func GetInstancel() *singletonl {
	if instancel == nil {
		instancel = new(singletonl)
	}
	return instancel
}

func (s *singletonl) doSomethingl() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstance()
	s.doSomething()
}
