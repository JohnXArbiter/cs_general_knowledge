package main

import "fmt"

func main() {
	factory := new(SimpleFactory)

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	grape := factory.CreateFruit("grape")
	grape.Show()
}

type sFruit interface {
	Show()
}

type sBanana struct {
}

func (*sBanana) Show() {
	fmt.Println("我是香蕉")
}

type sApple struct {
}

func (*sApple) Show() {
	fmt.Println("我是苹果")
}

type sGrape struct {
}

func (*sGrape) Show() {
	fmt.Println("我是葡萄")
}

type SimpleFactory struct {
}

func (*SimpleFactory) CreateFruit(kind string) sFruit {
	var fruit sFruit
	if kind == "apple" {
		fruit = new(sApple)
	} else if kind == "banana" {
		fruit = new(sBanana)
	} else if kind == "grape" {
		fruit = new(sGrape)
	}
	return fruit
}
