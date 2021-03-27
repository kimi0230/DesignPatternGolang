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
