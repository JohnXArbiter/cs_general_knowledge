package main

import "fmt"

func main() {
	//1. 制作一杯咖啡
	makeCoffee := NewMakeCoffee()
	makeCoffee.MakeBeverage() // 调用固定模板方法

	fmt.Println("------------")

	//2. 制作茶
	makeTea := NewMakeTea()
	makeTea.MakeBeverage()
}

// Beverage 抽象类，制作饮料,包裹一个模板的全部实现步骤
type Beverage interface {
	BoilWater()          // 煮开水
	Brew()               // 冲泡
	PourInCup()          // 倒入杯中
	AddThings()          // 添加酌料
	WantAddThings() bool // 是否加入酌料Hook
}

// 封装一套流程模板，让具体的制作流程继承且实现
type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()

	// 子类可以重写该方法来决定是否执行下面动作
	if t.b.WantAddThings() == true {
		t.b.AddThings()
	}
}

type MakeCoffee struct {
	template // 继承模板
}

func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)
	// b 为Beverage，是MakeCoffee的接口，这里需要给接口赋值，指向具体的子类对象
	// 来触发b全部接口方法的多态特性。
	makeCoffee.b = makeCoffee
	return makeCoffee
}

func (mc *MakeCoffee) BoilWater() {
	fmt.Println("将水煮到100摄氏度")
}

func (mc *MakeCoffee) Brew() {
	fmt.Println("用水冲咖啡豆")
}

func (mc *MakeCoffee) PourInCup() {
	fmt.Println("将充好的咖啡倒入陶瓷杯中")
}

func (mc *MakeCoffee) AddThings() {
	fmt.Println("添加牛奶和糖")
}

func (mc *MakeCoffee) WantAddThings() bool {
	return true //启动Hook条件
}

// MakeTea 具体的模板子类 制作茶
type MakeTea struct {
	template //继承模板
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	//b 为Beverage，是MakeTea，这里需要给接口赋值，指向具体的子类对象
	//来触发b全部接口方法的多态特性。
	makeTea.b = makeTea
	return makeTea
}

func (mt *MakeTea) BoilWater() {
	fmt.Println("将水煮到80摄氏度")
}

func (mt *MakeTea) Brew() {
	fmt.Println("用水冲茶叶")
}

func (mt *MakeTea) PourInCup() {
	fmt.Println("将充好的咖啡倒入茶壶中")
}

func (mt *MakeTea) AddThings() {
	fmt.Println("添加柠檬")
}

func (mt *MakeTea) WantAddThings() bool {
	return false //关闭Hook条件
}
