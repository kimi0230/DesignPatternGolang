### Observer 觀察者模式:
> 定義一對多的物件依存關係, 讓物件狀態一有變動, 就自動通知其他相依物件做該做的更新動作

定義了一種一對多的依賴關係, 讓多個觀察者物件同時監聽某一個主題物件. 
這主題物件在狀態發生變化時,會通知所有觀察者物件, 
使他們能夠自動更新自己

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Observer.png?raw=true)

### Example
https://github.com/kimi0230/DesignPatternGolang/tree/master/Observer 

```go
package observer

import "fmt"

// 抽象通知者 or 主題
type Subject interface {
	Attach(Observer)
	Detach(Observer)
	Notify()
	State() int
	SetState(int)
}

type ConcreteSubject struct {
	subjectState int
	observers    []*Observer
}

func (c *ConcreteSubject) Attach(o Observer) {
	if c == nil {
		return
	}
	for _, val := range c.observers {
		// 如果有訂閱 就不用在訂閱了
		if *val == o {
			fmt.Println("有訂閱 就不用在訂閱了")
			return
		}
	}
	c.observers = append(c.observers, &o)
}

func (c *ConcreteSubject) Detach(o Observer) {
	if c == nil {
		return
	}
	for key, val := range c.observers {
		// 有找到訂閱, 將其移除
		if *val == o {
			// fmt.Println("detach", (*val))
			c.observers = append(c.observers[:key], c.observers[key+1:]...)
		}
	}
}

func (c *ConcreteSubject) Notify() {
	if c == nil {
		return
	}
	for _, val := range c.observers {
		(*val).Update()
	}
}

func (c *ConcreteSubject) State() int {
	if c == nil {
		return 0
	}
	return c.subjectState
}

func (c *ConcreteSubject) SetState(state int) {
	if c == nil {
		return
	}
	c.subjectState = state
}

// 抽象觀察者
type Observer interface {
	Update()
}
type ConcreteObserver struct {
	name          string
	subject       Subject
	observerState int
}

func (c *ConcreteObserver) Update() {
	c.observerState = c.subject.State()
	fmt.Printf("觀察者 %s 的最新狀態是 %d \n", c.name, c.observerState)
}

func NewObserver(name string, subj Subject, state int) *ConcreteObserver {
	return &ConcreteObserver{name, subj, state}
}
```