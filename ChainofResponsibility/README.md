### Chain of Responsibility 職責鏈模式:
> 讓多個物件都有機會處理某一訊息, 以降低訊息發送者和接收者之間的耦合關係. 它將接收者物件串連起來, 讓訊息流經其中, 直到被處理了為止

使多個物件都有機會處理請求, 從而避免請求的發送者和接收者之間的耦合關係. 將這個物件連成一條鏈, 並沿著這條鏈傳遞該請求, 直到有一個物件處理它為止


![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/ChainofResponsibility.png?raw=true)

### Example
https://github.com/kimi0230/DesignPatternGolang/tree/master/ChainofResponsibility 

```go
package chainofresponsibility

import "fmt"

type IHandler interface {
	SetSuccessor(IHandler)
	HandleRequest(int)
}

type Handler struct {
	successor IHandler // 繼任者
}

// SetSuccessor : 設定繼任者
func (h *Handler) SetSuccessor(i IHandler) {
	if h == nil {
		return
	}
	h.successor = i
}

// 請求數在0~10之間則有權處理, 否則轉到下一位
type ConcreteHandler1 struct {
	Name string
	Handler
}

func NewConcreteHandler1(name string) *ConcreteHandler1 {
	return &ConcreteHandler1{name, Handler{}}
}

func (c *ConcreteHandler1) HandleRequest(req int) {
	if c == nil {
		return
	}
	if req >= 0 && req < 10 {
		fmt.Printf("%s 處理請求 %d \n", c.Name, req)
	} else if c.successor != nil {
		c.successor.HandleRequest(req)
	}
}

// 請求數在10~20之間則有權處理, 否則轉到下一位
type ConcreteHandler2 struct {
	Name string
	Handler
}

func NewConcreteHandler2(name string) *ConcreteHandler2 {
	return &ConcreteHandler2{name, Handler{}}
}

func (c *ConcreteHandler2) HandleRequest(req int) {
	if c == nil {
		return
	}
	if req >= 10 && req < 20 {
		fmt.Printf("%s 處理請求 %d \n", c.Name, req)
	} else if c.successor != nil {
		c.successor.HandleRequest(req)
	}
}

// 請求數在20~30之間則有權處理, 否則轉到下一位
type ConcreteHandler3 struct {
	Name string
	Handler
}

func NewConcreteHandler3(name string) *ConcreteHandler3 {
	return &ConcreteHandler3{name, Handler{}}
}

func (c *ConcreteHandler3) HandleRequest(req int) {
	if c == nil {
		return
	}
	if req >= 20 && req < 30 {
		fmt.Printf("%s 處理請求 %d \n", c.Name, req)
	} else if c.successor != nil {
		c.successor.HandleRequest(req)
	}
}
```