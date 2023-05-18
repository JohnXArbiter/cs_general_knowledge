package main

import "fmt"

func main() {
	g1 := goods{
		Kind: "韩国面膜",
		Fact: true,
	}

	g2 := goods{
		Kind: "CET4证书",
		Fact: false,
	}

	//如果不使用代理来完成从韩国购买任务
	var shopping Shopping
	shopping = new(korean) //具体的购买主题

	//1-先验货
	if g1.Fact == true {
		fmt.Println("对[", g1.Kind, "]进行了辨别真伪.")
		//2-去韩国购买
		shopping.buy(&g1)
		//3-海关安检
		fmt.Println("对[", g1.Kind, "] 进行了海关检查， 成功的带回祖国")
	}

	fmt.Println("---------------以下是 使用 代理模式-------")
	var overseasProxy Shopping
	overseasProxy = NewProxy(shopping)
	overseasProxy.buy(&g1)
	overseasProxy.buy(&g2)

	as := &goods{
		Kind: "烤肉",
		Fact: true,
	}

	overseasProxy.buy(as)
}

type goods struct {
	Kind string //商品种类
	Fact bool   //商品真伪
}

type Shopping interface {
	buy(goods *goods)
}

type korean struct {
}

func (*korean) buy(goods *goods) {
	fmt.Println("在韩国购物中，买了", goods.Kind)
}

type africa struct {
}

func (*africa) buy(goods *goods) {
	fmt.Println("在非洲购物中，买了", goods.Kind)
}

type america struct {
}

func (*america) buy(goods *goods) {
	fmt.Println("在阿美莉卡购物中，买了", goods.Kind)
}

type overseasProxy struct {
	Shopping
}

func (op *overseasProxy) buy(goods *goods) {
	// 1. 先验货
	if op.distinguish(goods) == true {
		//2. 进行购买
		op.Shopping.buy(goods) //调用原被代理的具体主题任务
		//3 海关安检
		op.check(goods)
	}
}

// 创建一个代理,并且配置关联被代理的主题
func NewProxy(shopping Shopping) Shopping {
	return &overseasProxy{shopping}
}

// 验货流程
func (op *overseasProxy) distinguish(goods *goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪.")
	if goods.Fact == false {
		fmt.Println("发现假货", goods.Kind, ", 不应该购买。")
	}
	return goods.Fact
}

// 安检流程
func (op *overseasProxy) check(goods *goods) {
	fmt.Println("对[", goods.Kind, "] 进行了海关检查， 成功的带回祖国")
}
