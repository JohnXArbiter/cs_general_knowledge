package main

import "fmt"

func main() {
	huawei := new(Huawei)
	kd := NewKeDecorator(huawei)
	//kd.ShowPhone()
	kmd := NewMoDecorator(kd)
	kmd.ShowPhone()
}

type DPhone interface {
	ShowPhone()
}

type Huawei struct {
}

func (*Huawei) ShowPhone() {
	fmt.Println("秀出了HuaWei手机")
}

type Xiaomian struct {
}

func (*Xiaomian) ShowPhone() {
	fmt.Println("秀出了XiaoMi手机")
}

type Decorator struct {
	DPhone
}

type MoDecorator struct {
	Decorator
}

type KeDecorator struct {
	Decorator
}

func (md *MoDecorator) ShowPhone() {
	md.DPhone.ShowPhone() //调用被装饰构件的原方法
	fmt.Println("贴膜的手机")  //装饰额外的方法
}

func NewMoDecorator(phone DPhone) DPhone {
	return &MoDecorator{Decorator{phone}}
}

func (kd *KeDecorator) ShowPhone() {
	kd.DPhone.ShowPhone() //调用被装饰构件的原方法
	fmt.Println("手机壳的手机") //装饰额外的方法
}

func NewKeDecorator(phone DPhone) DPhone {
	return &KeDecorator{Decorator{phone}}
}
