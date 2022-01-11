### Mediator 仲介者模式:
> 定義可將一群物件互動方式封裝起來的物件. 因為物件彼此不直接相互指涉, 所以耦合性低, 容易逐一變更互動關係

用一個仲介物件來封裝一系列的物件互動. 仲介者使各個物件不需要顯示地互相參考, 從而使其耦合鬆散, 而可以獨立低改變他們之間的互動

如果不存在擴展情況, Mediator 可以與 ConcreteMediator合二為一;
當系統出現了"多對多"互動複查的物件群時, 不要急著使用仲介者模式.
仲介者模式一般應用於一組物件以定義良好但是複雜的方式進行通訊的場合, 以及想訂製一個分布在多個類別中的行為, 而又不想產生太多子類別的場合

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Mediator.png?raw=true)

### Example
https://github.com/kimi0230/DesignPatternGolang/tree/master/Mediator 

```go
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

```