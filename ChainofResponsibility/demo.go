package chainofresponsibility

import "fmt"

type Request struct {
	requestType    string
	requestContent string
	number         int
}

func (r *Request) RequestType() string {
	if r == nil {
		return ""
	}
	return r.requestType
}
func (r *Request) SetRequestType(s string) {
	if r == nil {
		return
	}
	r.requestType = s
}

func (r *Request) RequestContent() string {
	if r == nil {
		return ""
	}
	return r.requestContent
}
func (r *Request) SetRequestContent(s string) {
	if r == nil {
		return
	}
	r.requestContent = s
}

func (r *Request) Number() int {
	if r == nil {
		return 0
	}
	return r.number
}
func (r *Request) SetNumber(i int) {
	if r == nil {
		return
	}
	r.number = i
}

// 管理者 抽象類別
type IManager interface {
	SetSuperior(IManager)
	RequestApplications(Request)
}

type Manager struct {
	Name     string
	superior IManager // 繼任者
}

func (m *Manager) SetSuperior(im IManager) {
	if m == nil {
		return
	}
	m.superior = im
}

// 經理 : ConcreteHandler
type CommonManager struct {
	Manager
}

func NewCommonManager(name string) *CommonManager {
	return &CommonManager{Manager{name, nil}}
}

func (c *CommonManager) RequestApplications(r Request) {
	if r.RequestType() == "請假" && r.Number() <= 2 {
		// 經理的權限就是可准許下屬請兩天內的假
		fmt.Printf("%s 批准 %s %d 天 \n", c.Name, r.RequestContent(), r.Number())
	} else {
		// 其餘的申請都需要轉到上級
		if c.superior != nil {
			c.superior.RequestApplications(r)
		}
	}
}

// 總監 : ConcreteHandler
type Director struct {
	Manager
}

func NewDirector(name string) *Director {
	return &Director{Manager{name, nil}}
}

func (c *Director) RequestApplications(r Request) {
	if r.RequestType() == "請假" && r.Number() <= 5 {
		// 總監的權限就是可准許下屬請一週
		fmt.Printf("%s 批准 %s, 數量 %d 天 \n", c.Name, r.RequestContent(), r.Number())
	} else {
		// 其餘的申請都需要轉到上級
		if c.superior != nil {
			c.superior.RequestApplications(r)
		}
	}
}

// 總經理 : ConcreteHandler
type GeneralManager struct {
	Manager
}

func NewGeneralManager(name string) *GeneralManager {
	return &GeneralManager{Manager{name, nil}}
}

func (c *GeneralManager) RequestApplications(r Request) {
	if r.RequestType() == "請假" {
		// 總經理的權限就是可准許下屬請假
		fmt.Printf("%s 批准 %s, 數量 %d 天  \n", c.Name, r.RequestContent(), r.Number())
	} else if r.RequestType() == "加薪" && r.Number() <= 500 {
		fmt.Printf("%s 批准 %s, 金額 %d  \n", c.Name, r.RequestContent(), r.Number())
	} else if r.RequestType() == "加薪" && r.Number() > 500 {
		fmt.Printf("%s: %s 加薪 %d ,請吃香蕉. 吱吱吱吱 \n", c.Name, r.RequestContent(), r.Number())
	} else {
		fmt.Printf("%s: %s 不知道 不清楚 我不會 \n", c.Name, r.RequestContent())
	}
}
