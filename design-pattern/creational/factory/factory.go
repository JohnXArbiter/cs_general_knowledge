package main

import "fmt"

func main() {
	bf := new(bananaFactory)
	banana := bf.createFruit()
	banana.Show()

	af := new(appleFactory)
	apple := af.createFruit()
	apple.Show()

	gf := new(grapeFactory)
	grape := gf.createFruit()
	grape.Show()
}

type Fruit interface {
	Show()
}

type Banana struct {
}

func (*Banana) Show() {
	fmt.Println("我是香蕉")
}

type Apple struct {
}

func (*Apple) Show() {
	fmt.Println("我是苹果")
}

type Grape struct {
}

func (*Grape) Show() {
	fmt.Println("我是葡萄")
}

type fruitFactory interface {
	createFruit()
}

type appleFactory struct {
}

func (*appleFactory) createFruit() Fruit {
	var apple Fruit
	apple = new(Apple)
	return apple
}

type bananaFactory struct {
}

func (*bananaFactory) createFruit() Fruit {
	var banana Fruit
	banana = new(Banana)
	return banana
}

type grapeFactory struct {
}

func (*grapeFactory) createFruit() Fruit {
	var grape Fruit
	grape = new(Grape)
	return grape
}
