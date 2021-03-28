package mediator

import "fmt"

// 定義一個抽象的發送消息方法, 得到同事物件和發送資訊
type IMediator interface {
	Send(string, IColleague)
}
type Mediator struct {
}

// 如果不存在擴展情況, Mediator 可以與 ConcreteMediator 合二為一
type ConcreteMediator struct {
	concreteColleague []IColleague
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{}
}

func (cm *ConcreteMediator) AddColleague(c IColleague) {
	if cm == nil {
		return
	}
	cm.concreteColleague = append(cm.concreteColleague, c)
}

func (cm *ConcreteMediator) Send(message string, colleague IColleague) {
	if cm == nil {
		return
	}
	for _, val := range cm.concreteColleague {
		if colleague == val {
			continue
		}
		val.Notify(message)
	}
}

// 抽象同事類別
type IColleague interface {
	Send(string)
	Notify(string)
}
type Colleague struct {
	mediator IMediator
}

type ConcreteColleague1 struct {
	Colleague
}

func NewConcreteColleague1(mediator IMediator) *ConcreteColleague1 {
	return &ConcreteColleague1{Colleague{mediator}}
}

func (c *ConcreteColleague1) Send(message string) {
	if c == nil {
		return
	}
	c.mediator.Send(message, c)
}

func (c *ConcreteColleague1) Notify(message string) {
	if c == nil {
		return
	}
	fmt.Println("同事1 得到訊息", message)
}

type ConcreteColleague2 struct {
	Colleague
}

func NewConcreteColleague2(mediator IMediator) *ConcreteColleague2 {
	return &ConcreteColleague2{Colleague{mediator}}
}

func (c *ConcreteColleague2) Send(message string) {
	if c == nil {
		return
	}
	c.mediator.Send(message, c)
}

func (c *ConcreteColleague2) Notify(message string) {
	if c == nil {
		return
	}
	fmt.Println("同事2 得到訊息", message)
}
