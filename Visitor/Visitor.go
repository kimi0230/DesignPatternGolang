package visitor

import "fmt"

type Visitor interface {
	VisitConcreteElementA(ConcreteElementA)
	VisitConcreteElementB(ConcreteElementB)
}

type ConcreteVisitorA struct {
}

func (c *ConcreteVisitorA) VisitConcreteElementA(ele ConcreteElementA) {
	if c == nil {
		return
	}
	ele.OperatorA()
}

func (c *ConcreteVisitorA) VisitConcreteElementB(ele ConcreteElementB) {
	if c == nil {
		return
	}
	ele.OperatorB()
}

type ConcreteVisitorB struct {
}

func (c *ConcreteVisitorB) VisitConcreteElementA(ele ConcreteElementA) {
	if c == nil {
		return
	}
	ele.OperatorA()
}

func (c *ConcreteVisitorB) VisitConcreteElementB(ele ConcreteElementB) {
	if c == nil {
		return
	}
	ele.OperatorB()
}

type Element interface {
	Accept(Visitor)
}

type ConcreteElementA struct{}

func (c *ConcreteElementA) Accept(visitor Visitor) {
	if c == nil {
		return
	}
	visitor.VisitConcreteElementA(*c)
}

func (c *ConcreteElementA) OperatorA() {
	if c == nil {
		return
	}
	fmt.Println("OperatorA")
}

type ConcreteElementB struct{}

func (c *ConcreteElementB) Accept(visitor Visitor) {
	if c == nil {
		return
	}
	visitor.VisitConcreteElementB(*c)
}

func (c *ConcreteElementB) OperatorB() {
	if c == nil {
		return
	}
	fmt.Println("OperatorB")
}

type ObjectStructure struct {
	elemetes []Element
}

func (o *ObjectStructure) Attach(ele Element) {
	if o == nil || ele == nil {
		return
	}
	o.elemetes = append(o.elemetes, ele)
}

func (o *ObjectStructure) Detach(ele Element) {
	if o == nil || ele == nil {
		return
	}
	for i, val := range o.elemetes {
		if val == ele {
			o.elemetes = append(o.elemetes[:i], o.elemetes[i+1:]...)
			break
		}
	}
}

func (o *ObjectStructure) Accept(v Visitor) {
	if o == nil {
		return
	}
	for _, val := range o.elemetes {
		val.Accept(v)
	}
}
