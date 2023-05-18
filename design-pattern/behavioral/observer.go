package main

import "fmt"

func main() {
	hero1 := Hero{
		"黄蓉",
		GaiBang,
	}

	hero2 := Hero{
		"洪七公",
		GaiBang,
	}

	hero3 := Hero{
		"乔峰",
		GaiBang,
	}

	hero4 := Hero{
		"张无忌",
		MingJiao,
	}

	hero5 := Hero{
		"韦一笑",
		MingJiao,
	}

	hero6 := Hero{
		"金毛狮王",
		MingJiao,
	}

	baixiao := BaiXiao{}

	baixiao.AddListener(&hero1)
	baixiao.AddListener(&hero2)
	baixiao.AddListener(&hero3)
	baixiao.AddListener(&hero4)
	baixiao.AddListener(&hero5)
	baixiao.AddListener(&hero6)

	fmt.Println("武林一片平静.....")
	hero1.Fight(&hero5, &baixiao)
}

const (
	GaiBang  = "丐帮"
	MingJiao = "明教"
)

// 抽象层
type Listener interface {
	OnFriendBeaten(event *Event)
	GetName() string
	GetParty() string
	Title() string
}

type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify(event *Event)
}

type Event struct {
	Notifier Notifier // 被知晓的通知者
	One      Listener // 事件主动发出者
	Another  Listener // 事件被动接收者
	Msg      string   // 具体的消息
}

// 实现层
type Hero struct {
	Name  string
	Party string
}

func (hero *Hero) Fight(another Listener, baixiao Notifier) {
	msg := fmt.Sprintf("%s 将 %s 揍了...", hero.Title(), another.Title())

	// 生成事件
	event := new(Event)
	event.Notifier = baixiao
	event.One = hero
	event.Another = another
	event.Msg = msg

	baixiao.Notify(event)
}

func (hero *Hero) Title() string {
	return fmt.Sprintf("[%s]:%s", hero.Party, hero.Name)
}

func (hero *Hero) OnFriendBeaten(event *Event) {
	// 判断是否为当事人
	if hero.Name == event.One.GetName() || hero.Name == event.Another.GetName() {
		return
	}

	// 本门派同伴将其他门派了，拍手叫好
	if hero.Party == event.One.GetParty() {
		fmt.Println(hero.Title(), "得知消息，拍手叫好！！！")
		return
	}

	// 本门派同伴被其他门派揍了，要主动报仇反击
	if hero.Party == event.Another.GetParty() {
		fmt.Println(hero.Title(), "得知消息，发起报仇反击！！！")
		hero.Fight(event.One, event.Notifier)
		return
	}
}

func (hero *Hero) GetName() string {
	return hero.Name
}

func (hero *Hero) GetParty() string {
	return hero.Party
}

// BaiXiao 百晓生(Nofifier)
type BaiXiao struct {
	heroList []Listener
}

// AddListener 添加观察者
func (b *BaiXiao) AddListener(listener Listener) {
	b.heroList = append(b.heroList, listener)
}

// RemoveListener 删除观察者
func (b *BaiXiao) RemoveListener(listener Listener) {
	for index, l := range b.heroList {
		//找到要删除的元素位置
		if listener == l {
			//将删除的点前后的元素链接起来
			b.heroList = append(b.heroList[:index], b.heroList[index+1:]...)
			break
		}
	}
}

func (b *BaiXiao) Notify(event *Event) {
	fmt.Println("【世界消息】 百晓生消息广播：", event.Msg)
	for _, listener := range b.heroList {
		listener.OnFriendBeaten(event)
	}
}
